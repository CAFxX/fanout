package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint }
func genFn0(p0 int32) int64 {
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint32 = 509
        v8 := GenS1{F0: 42}
        switch fi6 {
        case 446:
            v9 := GenS2{F0: ((((-2.25 + 3.14159) * 0.1) - (1e10 * 1.5)) >= 0.1)}
            _ = v9
        default:
            p0 = -506
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    var v10 int = 128
    var v11 int16 = 42
    if (42 == 131) {
        var v12 float32 = ((float32((((float64(-2.25) + ((((0.1 * float64(3.14159)) / float64(1.5)) / float64(3.14159)) / 0.1)) + (3.14159 * float64((1e10 - 1.5)))) * -2.25)) + ((1e10 + 1.5) + (0.0 / (0.1 - (0.0 * (3.14159 + -2.25)))))) + (float32(1.5) - (float32(3.14159) / -2.25)))
        var v13 float32 = v12
        for fi14 := 0; fi14 < 6; fi14++ {
            switch fi14 {
            case -430:
                v13 -= v13
            case 127:
                switch 322 {
                case 489:
                    if false {
                        v15 := []int{fi14, 127, (v10 << ((0) & 63)), 9223372036854775807}
                        v16 := GenS3{F0: 0}
                        var v17 bool = true
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                case 32767:
                    if (false || (-620 != -1)) {
                        v11 += 1000
                    } else {
                        v18 := []int{2147483647, (v10 * 798), v10}
                        _ = v18
                    }
                case 504:
                    v12 += ((v13 - 1.5) * ((1e10 * -2.25) * 0.0))
                }
            case 128:
                fmt.Println(fi14)
            }
            _ = fi14
        }
        _ = v12
        _ = v13
    } else {
        var v19 uint = 4294967295
        switch v10 {
        case 32768:
            for fi20 := 0; fi20 < 5; fi20++ {
                var v21 int16 = (v11 | (v11 << ((v19) & 63)))
                var v22 uint64 = uint64(v19)
                _ = fi20
                _ = v21
                _ = v22
            }
        case -2147483648:
            v23 := []int{1}
            _ = v23
        case 425:
            var v24 bool = ((false && false) && true)
            _ = v24
        default:
            v25 := GenS1{F0: -128}
            _ = v25
        }
        if false {
            v26 := []int{(v10 & -698), (v10 | 2147483647)}
            _ = v26
        }
        _ = v19
    }
    _ = p0
    _ = v10
    _ = v11
    return 1
}

func main() {
    var v1 uint32 = 686
    v2 := make(map[uint32]uint32)
    var v3 int16 = 478
    v4 := make(map[uint32]uint32)
    var v5 int32 = 43
    if true {
        for fi27 := 0; fi27 < 13; fi27++ {
            v28 := GenS3{F0: 65535}
            var v29 bool = (!("foo bar" <= ("a" + ("a" + ((("hello" + "a") + "a") + (("" + "") + "\n\t"))))))
            _ = fi27
            _ = v28
            _ = v29
        }
    } else {
        var v30 uint64 = 42
        for fi31 := 0; fi31 < 6; fi31++ {
            var v32 int64 = 70
            var v33 int32 = (v5 - (v5 - v5))
            v32 += (v32 >> ((uint(v1)) & 63))
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = v30
    }
    var v34 float64 = 0.0
    v35 := []int{-133, 2147483647, 729, -9223372036854775808}
    v5 = v5
    v36 := v2[v1]
    fmt.Println(v36)
    for fi37 := 0; fi37 < 1; fi37++ {
        v1 = v1
        _ = fi37
    }
    switch "a" {
    case "a":
        if (!(!false)) {
            switch ("" + "hello") {
            case "":
                for fi38 := 0; fi38 < 13; fi38++ {
                    var v39 uint32 = (v36 + 474)
                    _ = fi38
                    _ = v39
                }
            default:
                for fi40 := 0; fi40 < 17; fi40++ {
                    fmt.Println(v1, v34)
                    _ = fi40
                }
            }
            v41 := GenS2{F0: true}
            _ = v41
        }
    default:
        if len(v35) >= 2 {
            v35 = v35[1:len(v35)]
        }
    }
    v3 += v3
    v36 += v36
    if false {
        fmt.Println(v3)
        var v42 int = 9223372036854775807
        _ = v42
    }
    var v43 uint8 = uint8(v1)
    v36 += (v36 & 2)
    var v44 uint64 = 53
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v43
    _ = v44
    _ = v5
}
