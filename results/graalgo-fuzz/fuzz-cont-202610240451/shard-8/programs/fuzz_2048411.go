package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 65535
    var v3 float64 = 3.14159
    var v4 uint16 = 255
    var v5 float64 = 3.14159
    fmt.Println(v3, v3)
    v3 += ((-2.25 - v5) / (1.5 * float64(((float32(float64(0.1)) / (1e10 - 0.0)) - (0.0 + 1.5)))))
    if ((-143 < 1) && false) {
        v6 := GenS1{F0: 1}
        if (("hello" + "") != "foo bar") {
            for fi7 := 0; fi7 < 2; fi7++ {
                for fi8 := 0; fi8 < 2; fi8++ {
                    delete(v1, "\n\t")
                    _ = fi8
                }
                v9 := GenS1{F0: 68}
                _ = fi7
                _ = v9
            }
            var v10 int32 = 203
            v5 += 0.1
            _ = v10
        }
        v11 := []int{-1, 0, 1, 1}
        _ = v11
        _ = v6
    }
    if (1 <= 42) {
        v12 := []int{-1}
        switch -609 {
        case -1:
            for fi13 := 0; fi13 < 7; fi13++ {
                fmt.Println(fi13, v5, v3)
                _ = fi13
            }
        }
        _ = v12
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 int64 = -274
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 bool = (127 == -128)
            var v18 float64 = float64(-2.25)
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi14
        _ = v15
    }
    v2 = v2
    for fi19 := 0; fi19 < 5; fi19++ {
        if ((!((v5 - (1e10 / v3)) == ((float64(-2.25) * 0.1) - 3.14159))) || true) {
            fmt.Println(v2, v4)
            var v20 bool = true
            _ = v20
        } else {
            v21 := v1["hello"]
            v2 = uint32(v4)
            _ = v21
        }
        if false {
            v5 -= ((v3 - v3) - -2.25)
        }
        v22 := GenS1{F0: 1}
        _ = fi19
        _ = v22
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        switch "foo bar" {
        case "a":
            v4 = v4
        case "\n\t":
            v4 += 1
        }
        v5 += v3
        _ = fi23
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        v2 = 4294967295
        v1[("a" + "foo bar")] = (v2 & v2)
        switch 127 {
        case 2147483647:
            v3 -= ((3.14159 / float64(1.5)) / (float64(float32((0.1 + (v5 / -2.25)))) + (v5 / ((float64(3.14159) - ((v5 - v3) / 0.1)) + ((v5 * 1.5) + ((float64(-2.25) * (-2.25 + v5)) / ((v5 * 3.14159) * ((1.5 + float64(((-2.25 - (1e10 * float32(v5))) * float32(v5)))) - (v3 + (v5 - 3.14159))))))))))
        case 32767:
            if false {
                var v25 bool = (127 != 28)
                var v26 uint8 = 128
                if v25 {
                    for fi27 := 0; fi27 < 18; fi27++ {
                        fmt.Println(fi24, v26)
                        if (18446744073709551615 != 65535) {
                            v4 = 299
                            var v28 float64 = (v5 + float64((1.5 / 1e10)))
                            var v29 bool = ((false || (v25 && (-1 <= 100))) || (fi24 <= fi27))
                            v4 += v4
                            _ = v28
                            _ = v29
                        }
                        v1[(("\n\t" + ("" + "\n\t")) + "a")] = 844
                        _ = fi27
                    }
                    v30 := []int{590, (fi24 * (fi24 >> ((4294967295) & 63)))}
                    _ = v30
                } else {
                    for fi31 := 0; fi31 < 10; fi31++ {
                        v3 -= v5
                        v2 = v2
                        fmt.Println(v5, fi31, v5)
                        _ = fi31
                    }
                }
                _ = v25
                _ = v26
            } else {
                if (255 > 2) {
                    fmt.Println(v2, v2)
                }
                fmt.Println(v3, v4)
                switch (fi24 & -295) {
                case 32767:
                    v5 = 0.1
                case 127:
                    for fi32 := 0; fi32 < 7; fi32++ {
                        v3 -= float64(float32(v5))
                        v33 := GenS1{F0: 179}
                        var v34 float32 = -2.25
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                case 128:
                    var v35 uint64 = 1
                    _ = v35
                default:
                    for fi36 := 0; fi36 < 7; fi36++ {
                        switch fi36 {
                        case 4294967295:
                            v5 = v3
                        case 42:
                            v3 = float64((float32(v3) - 0.0))
                        }
                        _ = len(v1)
                        var v37 int16 = 32767
                        v5 += v3
                        _ = fi36
                        _ = v37
                    }
                }
            }
        default:
            v38 := []int{(fi24 << ((42) & 63)), 128, fi24}
            _ = v38
        }
        _ = fi24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
