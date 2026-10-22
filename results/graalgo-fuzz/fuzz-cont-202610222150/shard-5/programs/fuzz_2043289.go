package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int8 }
type GenS2 struct { F0 uint; F1 int64 }

func main() {
    var v1 uint8 = 128
    var v2 int16 = -206
    var v3 int16 = 1000
    v4 := []uint32{2, 2, 65535, 831}
    var v5 uint8 = 42
    v6 := make(map[string]uint32)
    v7 := GenS1{F0: v1, F1: 173, F2: 0}
    var v8 int8 = 100
    var v9 string = (("" + "\n\t") + ("hello" + "hello"))
    switch 42 {
    case -2147483648:
        if false {
            v8 -= (-(v8))
            var v10 uint32 = 42
            v11 := v6[v9]
            _ = v10
            _ = v11
        } else {
            if true {
                _ = len(v6)
                switch 4294967295 {
                case 0:
                    v12 := GenS1{F0: 2, F1: -618, F2: (v8 | v8)}
                    _ = v12
                case -1:
                    var v13 int = 32767
                    _ = v13
                case 32768:
                    var v14 uint64 = 2
                    _ = v14
                default:
                    switch 9223372036854775807 {
                    case 127:
                        var v15 bool = (-9223372036854775808 > 636)
                        _ = v15
                    case 605:
                        for fi16 := 0; fi16 < 5; fi16++ {
                            var v17 uint8 = v1
                            fmt.Println(v8, v2, v5)
                            fmt.Println(v17, fi16, v17)
                            _ = fi16
                            _ = v17
                        }
                    case 899:
                        switch (v9 + v9) {
                        case "":
                            for fi18 := 0; fi18 < 1; fi18++ {
                                v5 -= v1
                                for fi19 := 0; fi19 < 15; fi19++ {
                                    v1 += (^(v5))
                                    v9 = ((v9 + ("a" + ("a" + v9))) + v9)
                                    var v20 float64 = 1.5
                                    _ = fi19
                                    _ = v20
                                }
                                if false {
                                    for fi21 := 0; fi21 < 10; fi21++ {
                                        if (true || ((v2 / ((-32768) | 1)) <= v3)) {
                                            v8 = v8
                                            var v22 int32 = 42
                                            fmt.Println(v3, v5, v3)
                                            _ = v22
                                        }
                                        _ = fi21
                                    }
                                }
                                _ = fi18
                            }
                        case "\n\t":
                            if false {
                                var v23 uint = 42
                                v8 -= 4
                                _ = v23
                            }
                        default:
                            v24 := GenS2{F0: 42, F1: -628}
                            _ = v24
                        }
                    }
                }
            }
        }
    case 0:
        fmt.Println(v2, v1, v5)
    case 2147483647:
        var v25 uint16 = 2
        _ = v25
    }
    v9 = v9
    v26 := GenS2{F0: 255, F1: 9223372036854775807}
    var v27 uint32 = 1
    v5 = 54
    switch (v9 + v9) {
    case "foo bar":
        delete(v6, v9)
    }
    v28 := GenS2{F0: 2, F1: 0}
    v1 += (v1 - v1)
    delete(v6, (v9 + "hello"))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
