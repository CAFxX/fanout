package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string }
func genFn0(p0 int, p1 uint64, p2 int) uint32 {
    var v7 int8 = 103
    if (!((1 < 969) || true)) {
        for fi8 := 0; fi8 < 8; fi8++ {
            p2 -= (p2 + 32767)
            var v9 uint16 = 2
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 20; fi10++ {
            for fi11 := 0; fi11 < 17; fi11++ {
                v12 := []int{fi10, (p2 ^ (p2 | p0))}
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        switch (p2 & (fi13 | -1)) {
        case -9223372036854775808:
            v7 = v7
        case 4294967295:
            var v14 uint = 4294967295
            _ = v14
        }
        _ = fi13
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        p1 = 982
        var v16 uint64 = 42
        v17 := GenS1{F0: 42}
        _ = fi15
        _ = v16
        _ = v17
    }
    if false {
        if true {
            switch p2 {
            case -9223372036854775808:
                if (!true) {
                    if ((float32((-2.25 + 1e10)) == 0.1) || false) {
                        p1 -= (p1 - p1)
                        switch "a" {
                        case "foo bar":
                            switch p2 {
                            case 1:
                                var v18 string = "a"
                                _ = v18
                            }
                        case "a":
                            fmt.Println(p2)
                        case "hello":
                            var v19 uint = 18446744073709551615
                            _ = v19
                        default:
                            v7 -= v7
                        }
                        var v20 bool = false
                        _ = v20
                    }
                }
            }
            switch "\n\t" {
            case "a":
                p1 = 2
            case "hello":
                fmt.Println(p0, p2)
            case "\n\t":
                switch (("\n\t" + "hello") + ("hello" + "foo bar")) {
                case "\n\t":
                    if (!false) {
                        for fi21 := 0; fi21 < 12; fi21++ {
                            fmt.Println(p2, p2, v7)
                            for fi22 := 0; fi22 < 7; fi22++ {
                                var v23 bool = (-110 < v7)
                                switch "\n\t" {
                                case "":
                                    switch "hello" {
                                    case "\n\t":
                                        if false {
                                            v24 := []int{fi21, (fi21 << ((255) & 63))}
                                            _ = v24
                                        } else {
                                            var v25 uint32 = uint32(p1)
                                            var v26 int = (p2 % (((p2 + (p0 - (p2 << ((0) & 63))))) | 1))
                                            _ = v25
                                            _ = v26
                                        }
                                    case "":
                                        var v27 uint64 = (p1 | (p1 << ((912) & 63)))
                                        _ = v27
                                    case "a":
                                        if true {
                                            v23 = ((v23 || false) && (!(0 <= 0)))
                                            if ((1e10 != 1e10) || (!(0 >= -32768))) {
                                                fmt.Println(v23, p0, fi22)
                                                switch "" {
                                                case "hello":
                                                    if ((0.1 * (3.14159 + 3.14159)) >= (0.1 * 0.0)) {
                                                        fmt.Println(p1)
                                                        switch "hello" {
                                                        case "a":
                                                            var v28 int64 = int64(p2)
                                                            _ = v28
                                                        default:
                                                            var v29 string = ("" + "foo bar")
                                                            _ = v29
                                                        }
                                                    }
                                                case "foo bar":
                                                    var v30 uint64 = 4294967295
                                                    _ = v30
                                                default:
                                                    var v31 int32 = 968
                                                    v31 += v31
                                                    switch "a" {
                                                    case "foo bar":
                                                        if (v23 || v23) {
                                                            p0 = 972
                                                            switch "" {
                                                            case "foo bar":
                                                                fmt.Println(p1, fi21, p1)
                                                            default:
                                                                var v32 string = ("hello" + "\n\t")
                                                                _ = v32
                                                            }
                                                            v33 := GenS2{F0: 651}
                                                            _ = v33
                                                        }
                                                    case "":
                                                        v7 = -128
                                                    case "\n\t":
                                                        fmt.Println(v7, p1)
                                                    default:
                                                        v7 = v7
                                                    }
                                                    var v34 uint8 = 128
                                                    switch fi21 {
                                                    case 127:
                                                        v35 := []int{330, -507, 32767, fi22}
                                                        _ = v35
                                                    case -112:
                                                        if (0 == 32767) {
                                                            v36 := GenS3{F0: "foo bar"}
                                                            p2 = (^(fi22))
                                                            var v37 string = (("foo bar" + ((("\n\t" + "") + ("foo bar" + "a")) + "")) + "foo bar")
                                                            _ = v36
                                                            _ = v37
                                                        }
                                                    }
                                                    p2 += 42
                                                    _ = v31
                                                    _ = v34
                                                }
                                                v7 = v7
                                            }
                                        }
                                    default:
                                        fmt.Println(p0, v7)
                                    }
                                case "a":
                                    v23 = true
                                default:
                                    fmt.Println(p0)
                                }
                                _ = fi22
                                _ = v23
                            }
                            fmt.Println(v7)
                            _ = fi21
                        }
                    }
                case "":
                    v38 := GenS2{F0: -9223372036854775808}
                    _ = v38
                case "a":
                    p2 = p0
                }
            default:
                p2 -= 2147483647
            }
            var v39 bool = true
            _ = v39
        } else {
            var v40 uint64 = 1
            v40 = (^(p1))
            _ = v40
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2147483648
}

func main() {
    v1 := []int{}
    v2 := make(map[string]uint32)
    var v3 string = "\n\t"
    v4 := make(map[int]int)
    v5 := []uint32{4294967295}
    var v6 int16 = 985
    if false {
        var v41 int8 = 1
        _ = v41
    } else {
        fmt.Println(v3)
    }
    v42 := v4[128]
    if true {
        fmt.Println(v6, v42)
        if (!true) {
            for fi43 := 0; fi43 < 18; fi43++ {
                var v44 int16 = (v6 & (^(v6)))
                _ = fi43
                _ = v44
            }
            v42 = (v42 + (v42 - v42))
            switch v3 {
            case "foo bar":
                _ = len(v2)
            case "":
                if (-1 != 2147483647) {
                    v45 := GenS1{F0: 127}
                    delete(v4, (v42 << ((42) & 63)))
                    _ = v45
                } else {
                    if false {
                        var v46 int16 = (v6 | (v6 ^ 952))
                        _ = v46
                    }
                }
            }
        }
        var v47 int32 = 729
        _ = v47
    }
    switch v42 {
    case 42:
        if false {
            fmt.Println(v6, v42, v3)
            if (0 > 42) {
                _ = len(v4)
                v6 = v6
            }
        }
    default:
        fmt.Println(v42)
    }
    v1 = append(v1, v42)
    v48 := GenS2{F0: 9223372036854775807}
    for fi49 := 0; fi49 < 3; fi49++ {
        v4[fi49] = (fi49 >> ((2) & 63))
        v3 = v3
        _ = fi49
    }
    if false {
        v3 = v3
    } else {
        for fi50 := 0; fi50 < 11; fi50++ {
            v4[(v42 + 628)] = (fi50 & (v42 & v42))
            var v51 int8 = -128
            switch v3 {
            case "\n\t":
                var v52 uint = 1
                _ = v52
            default:
                if (!true) {
                    v42 = (v42 | 74)
                    switch (fi50 + (-(fi50))) {
                    case 1:
                        var v53 float64 = 1.5
                        _ = v53
                    default:
                        for fi54 := 0; fi54 < 14; fi54++ {
                            v1 = append(v1, 42)
                            _ = fi54
                        }
                    }
                    v55 := GenS1{F0: 1}
                    _ = v55
                } else {
                    switch (fi50 - fi50) {
                    case -145:
                        if (42 < 2147483648) {
                            var v56 uint8 = 42
                            fmt.Println(fi50, v3)
                            _ = v56
                        }
                    default:
                        v42 -= (v42 | fi50)
                    }
                }
            }
            _ = fi50
            _ = v51
        }
        v57 := GenS3{F0: "hello"}
        switch (v3 + ("hello" + "")) {
        case "\n\t":
            v58 := GenS4{F0: v3}
            _ = v58
        }
        _ = v57
    }
    v3 = v3
    switch (v42 << ((612) & 63)) {
    case 0:
        v59 := GenS3{F0: "hello"}
        _ = v59
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v48
    _ = v5
    _ = v6
}
