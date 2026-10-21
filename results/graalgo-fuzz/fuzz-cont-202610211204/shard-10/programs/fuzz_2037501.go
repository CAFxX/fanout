package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint, p1 uint64) string {
    fmt.Println(p0, p1)
    v6 := []int{}
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := GenS1{F0: p0}
        _ = fi7
        _ = v8
    }
    var v9 uint64 = p1
    if (!(true && (1 < 0))) {
        var v10 float32 = -2.25
        var v11 int64 = -478
        for fi12 := 0; fi12 < 12; fi12++ {
            for fi13 := 0; fi13 < 19; fi13++ {
                v14 := GenS1{F0: p0}
                var v15 int64 = -9223372036854775808
                switch fi12 {
                case 4294967295:
                    p1 -= v9
                case 1:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    switch "a" {
                    case "\n\t":
                        var v16 uint16 = 289
                        _ = v16
                    }
                    if false {
                        v6 = append(v6, (^(fi13)))
                        switch (^(fi12)) {
                        case 533:
                            var v17 int16 = 32767
                            _ = v17
                        }
                        var v18 int16 = -32768
                        _ = v18
                    }
                default:
                    if len(v6) > 0 {
                        v6[0] = 32768
                    }
                }
                _ = fi13
                _ = v14
                _ = v15
            }
            v10 += (v10 + -2.25)
            if false {
                v11 = (-(v11))
            }
            _ = fi12
        }
        _ = v10
        _ = v11
    } else {
        var v19 int8 = -47
        v19 += int8(p1)
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v9
    return "foo bar"
}

func main() {
    var v1 uint = 4294967295
    var v2 bool = (!false)
    var v3 int8 = 1
    v4 := make(map[uint32]uint32)
    var v5 float32 = 0.1
    v2 = (true && (v3 <= (v3 | (v3 - v3))))
    v20 := GenS1{F0: (v1 % (((^(v1))) | 1))}
    v1 = 37
    v21 := []int{985}
    var v22 int32 = -1
    for fi23 := 0; fi23 < 13; fi23++ {
        for fi24 := 0; fi24 < 10; fi24++ {
            v1 = (v1 - v1)
            switch fi24 {
            case 32768:
                var v25 int16 = 42
                _ = v25
            default:
                if len(v21) >= 2 {
                    v21 = v21[1:len(v21)]
                }
            }
            _ = fi24
        }
        v26 := GenS1{F0: v1}
        fmt.Println(fi23, v22, v2)
        _ = fi23
        _ = v26
    }
    var v27 int64 = 1
    var v28 string = ("foo bar" + "hello")
    if ((-9223372036854775808 < (v27 % ((v27) | 1))) && (((1.5 + (1.5 - -2.25)) / -2.25) <= 0.0)) {
        if (42 == 1000) {
            var v29 uint8 = 128
            _ = v29
        } else {
            if true {
                v21 = append(v21, 127)
                switch v28 {
                case "a":
                    var v30 uint16 = 935
                    _ = v30
                }
            }
        }
        v31 := GenS1{F0: (v1 / (((^(v1))) | 1))}
        _ = v31
    } else {
        v32 := GenS1{F0: (v1 + (v1 * 4294967295))}
        _ = v32
    }
    v5 = (((v5 + v5) * ((v5 / 1.5) * v5)) - v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
