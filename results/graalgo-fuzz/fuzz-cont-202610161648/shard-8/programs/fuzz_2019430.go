package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint16 }
func genFn0(p0 string, p1 bool) uint32 {
    if false {
        v6 := GenS1{F0: 0, F1: 85, F2: 42}
        fmt.Println(p1)
        _ = v6
    }
    v7 := []int{2147483647, 127}
    _ = p0
    _ = p1
    _ = v7
    return 2147483648
}
func genFn1(p0 bool, p1 bool) string {
    p1 = true
    p0 = p0
    var v8 int16 = 32767
    var v9 int32 = -1
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return ""
}

func main() {
    v1 := []int{-518, 32767}
    v2 := make(map[string]uint32)
    var v3 int16 = 556
    v4 := []int{32767, 4294967295, -1}
    var v5 uint64 = 1
    if false {
        switch "a" {
        case "\n\t":
            for fi10 := 0; fi10 < 6; fi10++ {
                v11 := GenS2{F0: (false && (!(!(v5 <= (v5 + 18446744073709551615))))), F1: 1}
                var v12 float64 = (0.1 / 1.0)
                v3 = int16(fi10)
                _ = fi10
                _ = v11
                _ = v12
            }
        }
    }
    var v13 uint = 1
    var v14 uint64 = v5
    var v15 uint64 = (v14 >> ((v13) & 63))
    if true {
        var v16 uint = (v13 - v13)
        _ = v16
    } else {
        var v17 int64 = -132
        if false {
            var v18 string = "a"
            _ = v18
        } else {
            var v19 int32 = int32(v17)
            _ = v19
        }
        if (!(true || ((v5 << ((v13) & 63)) == (v14 + (v15 + v5))))) {
            if true {
                var v20 string = "a"
                _ = v20
            }
        }
        _ = v17
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        fmt.Println(v15, fi21, v14)
        _ = fi21
    }
    switch "hello" {
    case "foo bar":
        v14 = v14
    case "":
        for fi22 := 0; fi22 < 20; fi22++ {
            if true {
                var v23 int32 = -1
                switch fi22 {
                case -943:
                    fmt.Println(v23)
                case -352:
                    fmt.Println(v14, v5)
                case 198:
                    var v24 bool = (-32768 <= (-(v3)))
                    _ = v24
                }
                _ = v23
            } else {
                v4 = append(v4, (^(fi22)))
            }
            v5 = 0
            _ = fi22
        }
    case "hello":
        v25 := GenS2{F0: true, F1: uint16(v15)}
        _ = v25
    default:
        if true {
            fmt.Println(v3)
        } else {
            v26 := GenS1{F0: v3, F1: -1, F2: 682}
            v13 = (v13 + 875)
            var v27 uint16 = 255
            _ = v26
            _ = v27
        }
    }
    v5 = v5
    var v28 bool = (!((!(!(-1 >= 2147483647))) || false))
    if len(v4) > 0 {
        v4[0] = -137
    }
    fmt.Println(v13, v15, v15)
    v29 := v2["\n\t"]
    v5 = (v5 / ((v14) | 1))
    var v30 int = -580
    v30 -= v30
    for fi31 := 0; fi31 < 1; fi31++ {
        var v32 uint8 = 2
        v33 := GenS1{F0: 42, F1: 100, F2: 529}
        if v28 {
            var v34 bool = ((v3 * v3) != v3)
            var v35 int8 = 122
            switch "a" {
            case "foo bar":
                v36 := v2[""]
                _ = v36
            case "\n\t":
                var v37 bool = (!false)
                _ = v37
            }
            _ = v34
            _ = v35
        } else {
            var v38 float32 = 1e10
            _ = v38
        }
        _ = fi31
        _ = v32
        _ = v33
    }
    for fi39 := 0; fi39 < 11; fi39++ {
        _ = len(v1)
        fmt.Println(v14, v3, v5)
        _ = fi39
    }
    if (v28 || (-1 != 42)) {
        if false {
            v15 = v14
            var v40 uint64 = (v14 >> ((v13) & 63))
            _ = v40
        }
        var v41 uint64 = v15
        _ = v41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
