package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 int }
func genFn0() uint32 {
    var v7 int32 = -1
    v7 -= v7
    v8 := []int{42, 473}
    _ = v7
    _ = v8
    return 2147483648
}
func genFn1(p0 int) bool {
    fmt.Println(p0)
    for fi9 := 0; fi9 < 2; fi9++ {
        p0 = (-(fi9))
        v10 := []int{32767, (p0 | (fi9 - -2147483648)), p0}
        _ = fi9
        _ = v10
    }
    v11 := []int{p0, p0, 694, p0}
    _ = p0
    _ = v11
    return false
}

func main() {
    var v1 int64 = 1
    var v2 bool = false
    var v3 int64 = v1
    v4 := make(map[string]string)
    v5 := make(map[int]uint32)
    v6 := make(map[int]int)
    v1 += 9223372036854775807
    v12 := GenS1{F0: 196}
    var v13 uint16 = 65535
    fmt.Println(v2, v3, v13)
    for fi14 := 0; fi14 < 18; fi14++ {
        for fi15 := 0; fi15 < 15; fi15++ {
            var v16 uint16 = v13
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    v17 := []int{127, -2147483648, 127}
    for fi18 := 0; fi18 < 13; fi18++ {
        v19 := GenS2{F0: 2147483647}
        v3 += (v3 << ((915) & 63))
        v13 -= v13
        _ = fi18
        _ = v19
    }
    if len(v17) > 0 {
        v17[0] = 127
    }
    v20 := GenS1{F0: 128}
    var v21 uint = 2
    for fi22 := 0; fi22 < 17; fi22++ {
        var v23 float32 = 1e10
        _ = fi22
        _ = v23
    }
    switch 2147483647 {
    case -1:
        if v2 {
            v3 = v1
            switch -1 {
            case 32767:
                if v2 {
                    var v24 int64 = (v1 & 0)
                    v2 = v2
                    switch ((("\n\t" + ("" + ("foo bar" + ("hello" + "\n\t")))) + "\n\t") + "a") {
                    case "\n\t":
                        switch -2147483648 {
                        case 2147483647:
                            v24 = (v3 + v1)
                        case 474:
                            _ = len(v17)
                        case 32767:
                            v25 := v5[4294967295]
                            _ = v25
                        }
                    }
                    _ = v24
                } else {
                    var v26 uint16 = (^(v13))
                    v27 := GenS2{F0: 1}
                    _ = v26
                    _ = v27
                }
            case 230:
                v2 = v2
            case 4294967295:
                for fi28 := 0; fi28 < 14; fi28++ {
                    v21 = v21
                    _ = fi28
                }
            }
        }
    case 2147483647:
        var v29 int = 580
        _ = v29
    case -170:
        v3 = 9223372036854775807
    default:
        var v30 uint16 = 803
        _ = v30
    }
    var v31 uint8 = 0
    v2 = v2
    var v32 bool = (4294967295 >= 561)
    fmt.Println(v2)
    v33 := GenS3{F0: 0, F1: 2147483647}
    _ = len(v17)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
