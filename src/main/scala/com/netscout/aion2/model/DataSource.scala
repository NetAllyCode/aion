package com.netscout.aion2.model

trait DataSource {
  def classOfType(t: String): Class[_]
  def initializeSchema(objects: Set[AionObjectConfig])
  def insertQuery(obj: AionObjectConfig, index: AionIndexConfig, values: Map[String, AnyRef], splitKeyValue: AnyRef)
  def executeQuery(obj: AionObjectConfig, index: AionIndexConfig, query: QueryStrategy, partitionConstraints: Map[String, AnyRef]): Iterable[Iterable[(String, Object)]]
}