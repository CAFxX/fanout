package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int32, p1 int8) bool {
    var v4 int = 32767
    if false {
        v5 := []int{}
        _ = v5
    }
    _ = p0
    _ = p1
    _ = v4
    return ((((0.1 + 3.14159) + (1e10 / 3.14159)) * (float64(0.0) * 0.0)) < 0.0)
}

func main() {
    var v1 int = 655
    var v2 uint64 = 18446744073709551615
    v3 := make(map[string]bool)
    var v6 uint64 = v2
    if (false && ((!false) || false)) {
        var v7 uint64 = v6
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 int8 = 100
            var v10 uint = 25
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v7
    }
    if true {
        if true {
            v1 = v1
            v2 = v6
            var v11 float32 = ((1e10 / 3.14159) - 1.5)
            _ = v11
        }
    } else {
        for fi12 := 0; fi12 < 8; fi12++ {
            if (!(1.5 <= 1.5)) {
                v1 += v1
            } else {
                v6 = v6
                var v13 bool = false
                _ = v13
            }
            if true {
                var v14 uint8 = 126
                _ = v14
            }
            v15 := GenS1{F0: 2}
            _ = fi12
            _ = v15
        }
        if true {
            v2 = v6
        }
        v1 -= v1
    }
    _ = len(v3)
    v16 := GenS1{F0: 2}
    fmt.Println(v1)
    v2 = v2
    if false {
        for fi17 := 0; fi17 < 5; fi17++ {
            for fi18 := 0; fi18 < 11; fi18++ {
                fmt.Println(fi18)
                switch v1 {
                case 128:
                    switch "hello" {
                    case "\n\t":
                        var v19 uint64 = (^(v6))
                        _ = v19
                    default:
                        if false {
                            v6 = 18446744073709551615
                            v20 := GenS1{F0: 1}
                            _ = v20
                        }
                    }
                case 752:
                    if false {
                        v6 = (v6 - 4294967295)
                        switch fi18 {
                        case 2147483647:
                            v2 = (v2 << ((18446744073709551615) & 63))
                        case -233:
                            switch (fi17 & v1) {
                            case 128:
                                switch (fi18 + (-(fi18))) {
                                case -990:
                                    v1 -= (fi18 - (fi18 >> ((1) & 63)))
                                }
                            }
                        default:
                            v2 -= (v6 >> ((4294967295) & 63))
                        }
                    }
                case 970:
                    var v21 int64 = 564
                    _ = v21
                default:
                    v22 := []int{102, -1, (v1 / (((fi18 * fi18)) | 1))}
                    _ = v22
                }
                v6 -= v6
                _ = fi18
            }
            _ = len(v3)
            _ = fi17
        }
        if false {
            var v23 bool = true
            _ = v23
        } else {
            v3["hello"] = (426 == 1)
        }
    }
    if false {
        v6 += (v2 << ((42) & 63))
        v24 := []int{v1}
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v6
}
