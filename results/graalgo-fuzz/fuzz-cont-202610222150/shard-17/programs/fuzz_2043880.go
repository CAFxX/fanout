package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int64, p1 int16, p2 bool) int8 {
    p2 = ((!p2) && false)
    switch "foo bar" {
    case "":
        fmt.Println(p0, p2)
    case "foo bar":
        p0 += 9223372036854775807
    default:
        var v7 uint64 = 2
        _ = v7
    }
    var v8 float64 = (-2.25 - 0.0)
    v9 := []int{763, 32767}
    var v10 uint32 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 100
}

func main() {
    var v1 uint = 4294967295
    var v2 int = 32767
    var v3 uint32 = 65535
    v4 := []string{"a", "", (("foo bar" + (("\n\t" + "\n\t") + ("\n\t" + "foo bar"))) + "foo bar"), "a"}
    var v5 float64 = (3.14159 * (float64(1.5) - float64(0.1)))
    v6 := make(map[int]uint32)
    v11 := GenS1{F0: -2147483648}
    for fi12 := 0; fi12 < 17; fi12++ {
        delete(v6, fi12)
        var v13 int64 = 1
        for fi14 := 0; fi14 < 4; fi14++ {
            if true {
                v15 := GenS1{F0: -2147483648}
                _ = v15
            }
            v5 = 0.1
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    var v16 float64 = v5
    v3 = v3
    var v17 uint = (^(v1))
    switch (v2 << ((v1) & 63)) {
    case -251:
        if true {
            var v18 uint16 = 32767
            _ = len(v6)
            _ = v18
        } else {
            var v19 float64 = v16
            for fi20 := 0; fi20 < 9; fi20++ {
                v21 := GenS2{F0: v3}
                for fi22 := 0; fi22 < 9; fi22++ {
                    var v23 uint8 = uint8(v1)
                    var v24 uint8 = (^(v23))
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                _ = fi20
                _ = v21
            }
            for fi25 := 0; fi25 < 20; fi25++ {
                v4 = append(v4, "\n\t")
                v17 -= v1
                switch 797 {
                case 589:
                    delete(v6, v2)
                case 9223372036854775807:
                    for fi26 := 0; fi26 < 12; fi26++ {
                        fmt.Println(v17, v5)
                        if len(v4) > 0 {
                            v4[0] = "foo bar"
                        }
                        v16 -= 0.0
                        if false {
                            if len(v4) > 0 {
                                v4[0] = "foo bar"
                            }
                            v5 += 1e10
                        }
                        v16 = v16
                        if false {
                            v19 = (v5 * (v5 * v5))
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        _ = fi26
                    }
                default:
                    delete(v6, -394)
                }
                _ = fi25
            }
            _ = v19
        }
    case 128:
        var v27 uint32 = 502
        _ = v27
    case 759:
        v1 -= v17
    default:
        var v28 int8 = -1
        _ = v28
    }
    fmt.Println(v1, v3)
    v5 = v16
    var v29 uint8 = 255
    v30 := GenS1{F0: 2147483647}
    for fi31 := 0; fi31 < 10; fi31++ {
        if true {
            v17 += 255
            var v32 int8 = 90
            _ = v32
        }
        for fi33 := 0; fi33 < 8; fi33++ {
            if (v17 == (v1 - v17)) {
                var v34 int32 = 204
                var v35 uint64 = 2
                _ = v34
                _ = v35
            }
            var v36 uint16 = 42
            fmt.Println(v3)
            _ = fi33
            _ = v36
        }
        _ = fi31
    }
    if len(v4) > 0 {
        v4[0] = (((("foo bar" + "\n\t") + ("hello" + "\n\t")) + ("a" + "\n\t")) + "a")
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
