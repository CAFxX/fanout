package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint64; F1 int64; F2 uint64 }

func main() {
    var v1 int32 = -349
    var v2 uint8 = 1
    var v3 float64 = (0.1 - 0.0)
    v4 := make(map[string]uint32)
    var v5 uint32 = 2
    var v6 uint32 = v5
    delete(v4, ("\n\t" + ""))
    var v7 int16 = 1000
    v8 := []int{399, 584}
    switch "" {
    case "":
        if ((!(v1 <= v1)) && (!((v3 * ((v3 * 0.1) / 1.5)) >= 1.5))) {
            for fi9 := 0; fi9 < 3; fi9++ {
                for fi10 := 0; fi10 < 13; fi10++ {
                    var v11 int16 = (v7 << ((481) & 63))
                    _ = fi10
                    _ = v11
                }
                if false {
                    if true {
                        v3 -= (v3 * ((((v3 * (v3 + 0.1)) + 1e10) * v3) / v3))
                    }
                    if (!true) {
                        for fi12 := 0; fi12 < 11; fi12++ {
                            if (!true) {
                                v13 := GenS1{F0: (v6 * uint32(v2)), F1: 9223372036854775807}
                                _ = len(v4)
                                var v14 uint64 = 2
                                _ = v13
                                _ = v14
                            }
                            _ = fi12
                        }
                        for fi15 := 0; fi15 < 14; fi15++ {
                            var v16 int8 = 1
                            switch "hello" {
                            case "\n\t":
                                delete(v4, "hello")
                            case "foo bar":
                                fmt.Println(v16, v2)
                            case "":
                                fmt.Println(v5, v5, v1)
                            }
                            _ = fi15
                            _ = v16
                        }
                    } else {
                        var v17 uint = 65535
                        _ = len(v4)
                        var v18 float32 = 0.1
                        _ = v17
                        _ = v18
                    }
                }
                _ = fi9
            }
        }
    }
    fmt.Println(v2, v7, v6)
    var v19 uint16 = 1
    v20 := GenS2{F0: 1, F1: -172, F2: 303}
    fmt.Println(v19)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
