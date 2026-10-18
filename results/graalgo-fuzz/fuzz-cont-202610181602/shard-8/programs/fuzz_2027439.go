package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
func genFn0() uint32 {
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint64 = 981
        fmt.Println(v7, fi6)
        _ = fi6
        _ = v7
    }
    switch 330 {
    case -724:
        v8 := GenS1{F0: -1}
        _ = v8
    case -1:
        var v9 uint32 = 4294967295
        _ = v9
    }
    v10 := []int{127}
    _ = v10
    return 473
}

func main() {
    v1 := make(map[int]string)
    var v2 string = (("a" + ("foo bar" + "a")) + "\n\t")
    var v3 uint32 = 2
    var v4 int16 = 32767
    var v5 uint16 = 775
    switch 808 {
    case 9223372036854775807:
        if false {
            switch v2 {
            case "hello":
                v11 := []int{-123}
                _ = v11
            }
        }
    case -2147483648:
        v2 = "foo bar"
    default:
        var v12 int64 = 0
        _ = v12
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 float32 = ((-2.25 * 0.1) + (float32(1e10) * 1e10))
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 int32 = -286
        v17 := GenS1{F0: 100}
        _ = fi15
        _ = v16
        _ = v17
    }
    v5 += 65535
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 int = 698
        _ = fi18
        _ = v19
    }
    v20 := []int{1, -9223372036854775808, 9223372036854775807}
    v5 -= v5
    for fi21 := 0; fi21 < 18; fi21++ {
        for fi22 := 0; fi22 < 12; fi22++ {
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            v2 = v2
            v23 := GenS2{F0: (v4 << ((42) & 63))}
            switch (fi22 % ((fi21) | 1)) {
            case -9223372036854775808:
                switch fi22 {
                case 127:
                    fmt.Println(fi21)
                default:
                    if false {
                        if true {
                            var v24 int16 = 42
                            v25 := v1[(-(fi21))]
                            _ = v24
                            _ = v25
                        } else {
                            var v26 uint16 = v5
                            v26 += uint16(v3)
                            v2 = ("\n\t" + "\n\t")
                            _ = v26
                        }
                        v20 = append(v20, (fi21 ^ (fi22 - (fi21 - fi21))))
                    }
                }
            default:
                v27 := GenS2{F0: -32768}
                _ = v27
            }
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    v5 = (^(v5))
    v2 = v2
    v20 = append(v20, -9223372036854775808)
    fmt.Println(v3, v5)
    v28 := GenS1{F0: 62}
    for fi29 := 0; fi29 < 2; fi29++ {
        if true {
            fmt.Println(fi29, v5)
            if true {
                v20 = append(v20, (fi29 * (fi29 + (fi29 + (fi29 - (fi29 ^ (^(fi29))))))))
                if (100 > 9) {
                    switch (v2 + v2) {
                    case "foo bar":
                        if len(v20) >= 2 {
                            v20 = v20[1:len(v20)]
                        }
                    case "":
                        v30 := v1[(fi29 / ((fi29) | 1))]
                        _ = v30
                    case "a":
                        var v31 int16 = (v4 << ((42) & 63))
                        _ = v31
                    default:
                        v4 = v4
                    }
                }
            }
        } else {
            if true {
                var v32 float32 = 1e10
                switch (("" + v2) + v2) {
                case "":
                    v20 = append(v20, (fi29 << ((207) & 63)))
                default:
                    v33 := v1[fi29]
                    _ = v33
                }
                _ = v32
            } else {
                v20 = append(v20, fi29)
            }
            v5 = 1
            var v34 uint64 = 0
            _ = v34
        }
        _ = fi29
    }
    v3 = (v3 % ((106) | 1))
    v35 := GenS2{F0: 1000}
    v3 -= v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
