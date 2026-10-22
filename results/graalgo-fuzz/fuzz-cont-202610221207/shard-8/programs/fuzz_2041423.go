package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{}
    v3 := make(map[string]uint32)
    var v4 int16 = -1
    var v5 float64 = float64(-2.25)
    v5 -= (v5 - v5)
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 int32 = -677
        if ((fi6 - fi6) <= (fi6 << ((977) & 63))) {
            var v8 uint = 4294967295
            _ = v8
        }
        v1[778] = "foo bar"
        _ = fi6
        _ = v7
    }
    v9 := GenS1{F0: "a"}
    v5 = v5
    v5 += v5
    v10 := v3["hello"]
    var v11 int16 = (v4 >> ((454) & 63))
    var v12 float32 = 3.14159
    var v13 bool = true
    if false {
        v14 := GenS2{F0: -86}
        _ = v14
    } else {
        v12 = 3.14159
        v11 += 1000
    }
    var v15 uint = 657
    switch (("a" + ((("" + "a") + "a") + "")) + (("hello" + "\n\t") + "hello")) {
    case "a":
        _ = len(v2)
    case "":
        switch 127 {
        case 580:
            switch 0 {
            case -904:
                for fi16 := 0; fi16 < 10; fi16++ {
                    v17 := GenS1{F0: "\n\t"}
                    _ = fi16
                    _ = v17
                }
            case 4294967295:
                switch 4294967295 {
                case 9223372036854775807:
                    for fi18 := 0; fi18 < 1; fi18++ {
                        if v13 {
                            if true {
                                var v19 int8 = 1
                                fmt.Println(v4, v5, v4)
                                _ = v19
                            }
                            v20 := GenS1{F0: ""}
                            for fi21 := 0; fi21 < 10; fi21++ {
                                var v22 uint64 = 2
                                _ = fi21
                                _ = v22
                            }
                            _ = v20
                        } else {
                            v12 -= v12
                        }
                        v23 := GenS1{F0: ""}
                        var v24 int16 = 32767
                        _ = fi18
                        _ = v23
                        _ = v24
                    }
                case 976:
                    for fi25 := 0; fi25 < 10; fi25++ {
                        for fi26 := 0; fi26 < 2; fi26++ {
                            v15 += v15
                            _ = fi26
                        }
                        fmt.Println(v5, v11)
                        _ = fi25
                    }
                case 1:
                    var v27 uint = 255
                    _ = v27
                default:
                    delete(v1, (^(v10)))
                }
            default:
                var v28 int = 32767
                _ = v28
            }
        case 1:
            fmt.Println(v10, v13)
        default:
            fmt.Println(v4)
        }
    case "\n\t":
        for fi29 := 0; fi29 < 14; fi29++ {
            if (!false) {
                delete(v3, "foo bar")
                v2 = append(v2, (v10 - v10))
            }
            var v30 float64 = 1e10
            _ = fi29
            _ = v30
        }
    }
    v12 = (3.14159 * (1.5 - 3.14159))
    v4 -= v11
    var v31 uint64 = 1
    var v32 int = -501
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v9
}
