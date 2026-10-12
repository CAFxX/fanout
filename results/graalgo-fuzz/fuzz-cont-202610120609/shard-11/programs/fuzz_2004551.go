package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 uint32 = 599
    v2 := []uint32{v1, v1}
    var v3 int64 = 121
    var v4 int64 = 0
    var v5 int32 = 42
    v6 := make(map[string]string)
    v6["a"] = ("foo bar" + "\n\t")
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := GenS1{F0: 0}
        switch "a" {
        case "foo bar":
            if (175 >= 922) {
                if true {
                    if len(v2) > 0 {
                        v2[0] = (v1 + v1)
                    }
                }
            }
        }
        for fi9 := 0; fi9 < 6; fi9++ {
            v6["hello"] = "\n\t"
            if false {
                v1 = v1
                v10 := GenS1{F0: 1}
                _ = v10
            } else {
                fmt.Println(v4)
                var v11 int = fi9
                if len(v2) > 0 {
                    v2[0] = (v1 << ((1) & 63))
                }
                _ = v11
            }
            if false {
                if true {
                    if (1 < 714) {
                        var v12 uint32 = v1
                        fmt.Println(v12)
                        if (42 <= 65535) {
                            fmt.Println(v4, v4, v5)
                            v3 = v4
                            v1 += (v1 / ((v12) | 1))
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v13 int64 = 0
                            _ = v13
                        }
                        _ = v12
                    }
                } else {
                    v14 := GenS1{F0: 1}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v15 := GenS1{F0: 42}
                    _ = v14
                    _ = v15
                }
            }
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 uint16 = 42
        _ = fi16
        _ = v17
    }
    v3 = (v3 << ((4294967295) & 63))
    v4 = (v4 / (((-(v4))) | 1))
    fmt.Println(v4)
    for fi18 := 0; fi18 < 9; fi18++ {
        for fi19 := 0; fi19 < 8; fi19++ {
            v2 = append(v2, (v1 << ((42) & 63)))
            v20 := GenS1{F0: 2}
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 1; fi21++ {
            v22 := v6[("\n\t" + "a")]
            var v23 int16 = -1
            var v24 int = fi21
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        _ = fi18
    }
    if (!true) {
        var v25 uint = 2
        var v26 uint32 = 42
        v4 += v3
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
