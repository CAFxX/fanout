package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = 86
    v3 := make(map[string]string)
    var v4 uint64 = 2
    switch ("\n\t" + "a") {
    case "\n\t":
        for fi5 := 0; fi5 < 15; fi5++ {
            v6 := v3[(("foo bar" + ("hello" + "foo bar")) + "\n\t")]
            if true {
                for fi7 := 0; fi7 < 16; fi7++ {
                    var v8 uint64 = 2
                    fmt.Println(fi5, fi5)
                    if (!(42 < v8)) {
                        v9 := GenS1{F0: -920}
                        var v10 string = v6
                        switch fi7 {
                        case 9223372036854775807:
                            v8 -= (v4 - 2)
                        case 4294967295:
                            fmt.Println(fi7, fi5, v6)
                        case 458:
                            switch -541 {
                            case 0:
                                v2 = v2
                            case 42:
                                var v11 int64 = 1
                                _ = v11
                            default:
                                v8 = (v4 / (((v4 ^ v8)) | 1))
                            }
                        default:
                            var v12 uint64 = 0
                            _ = v12
                        }
                        _ = v10
                        _ = v9
                    } else {
                        var v13 string = v6
                        fmt.Println(v2)
                        _ = v13
                    }
                    _ = fi7
                    _ = v8
                }
            }
            _ = fi5
            _ = v6
        }
    case "hello":
        var v14 bool = (false || (-2.25 != ((-2.25 - 0.1) + ((1e10 * 0.0) * 0.0))))
        _ = v14
    }
    fmt.Println(v4)
    v2 = v2
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 uint32 = 84
        _ = fi15
        _ = v16
    }
    v17 := GenS1{F0: -9223372036854775808}
    v4 = (v4 % (((^(v4))) | 1))
    switch "foo bar" {
    case "a":
        for fi18 := 0; fi18 < 7; fi18++ {
            v4 = 0
            v3[""] = "a"
            v19 := []int{(^(fi18)), fi18, (fi18 << ((2) & 63))}
            _ = fi18
            _ = v19
        }
    default:
        var v20 int64 = 42
        _ = v20
    }
    fmt.Println(v4, v4)
    v21 := GenS1{F0: 1}
    v4 = (^(v4))
    var v22 uint32 = 4294967295
    v1[(^(v22))] = v22
    v22 = 42
    v4 -= (v4 | (v4 | v4))
    var v23 int64 = 871
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
