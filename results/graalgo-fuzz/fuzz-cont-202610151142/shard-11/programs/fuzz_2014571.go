package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint64; F1 int32 }
func genFn0() string {
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := []int{(-(fi7)), 453, (fi7 << ((767) & 63)), (fi7 - fi7)}
        if (0.1 == -2.25) {
            var v9 uint32 = 42
            for fi10 := 0; fi10 < 15; fi10++ {
                var v11 bool = (0 != 1)
                v9 = v9
                _ = len(v8)
                _ = fi10
                _ = v11
            }
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        fmt.Println(fi12)
        var v13 int32 = 2147483647
        v13 -= 0
        _ = fi12
        _ = v13
    }
    return ("foo bar" + ("" + "hello"))
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{"\n\t", (("a" + ("a" + "\n\t")) + "\n\t"), "hello", "a"}
    var v3 int = -1
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]uint32)
    var v6 uint = 237
    var v14 bool = true
    var v15 int16 = 1000
    v2 = append(v2, ("" + ""))
    var v16 string = "a"
    var v17 uint64 = 18446744073709551615
    if true {
        var v18 uint64 = v17
        for fi19 := 0; fi19 < 12; fi19++ {
            var v20 uint16 = 2
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 13; fi21++ {
            v22 := GenS1{F0: (18446744073709551615 == v17), F1: false, F2: v18}
            for fi23 := 0; fi23 < 11; fi23++ {
                v6 += (v6 - (v6 / (((^(v6))) | 1)))
                v1[(fi23 ^ 32767)] = ("a" + (((v16 + (v16 + "\n\t")) + v16) + (v16 + "a")))
                var v24 float64 = (0.1 - -2.25)
                _ = fi23
                _ = v24
            }
            switch v16 {
            case "foo bar":
                switch "foo bar" {
                case "hello":
                    fmt.Println(v16)
                case "\n\t":
                    var v25 uint64 = 506
                    _ = v25
                default:
                    fmt.Println(v3, v3)
                }
            case "a":
                for fi26 := 0; fi26 < 11; fi26++ {
                    v27 := GenS1{F0: ((!((752 == 928) && true)) || v14), F1: (v14 && (v14 || (("foo bar" + (v16 + (((v16 + v16) + v16) + (v16 + ("a" + (v16 + (v16 + v16))))))) == v16))), F2: 4294967295}
                    _ = fi26
                    _ = v27
                }
            default:
                var v28 int16 = v15
                _ = v28
            }
            _ = fi21
            _ = v22
        }
        _ = v18
    }
    var v29 int32 = int32(v15)
    if false {
        switch (v16 + (((v16 + v16) + (v16 + v16)) + v16)) {
        case "":
            v30 := GenS2{F0: (^(v17)), F1: v29}
            _ = v30
        default:
            if false {
                var v31 uint = v6
                v32 := GenS2{F0: 42, F1: (^(v29))}
                _ = v31
                _ = v32
            } else {
                var v33 int16 = v15
                fmt.Println(v29, v14, v16)
                var v34 uint16 = 32767
                _ = v33
                _ = v34
            }
        }
        v14 = true
    }
    var v35 int16 = v15
    v36 := GenS2{F0: (^(v17)), F1: 957}
    var v37 string = (((("a" + "\n\t") + v16) + v16) + v16)
    var v38 int = -862
    switch "foo bar" {
    case "":
        if true {
            v39 := GenS2{F0: (v17 | 1), F1: -1}
            _ = v39
        }
    case "\n\t":
        switch 128 {
        case 4294967295:
            var v40 int8 = 4
            _ = v40
        }
    case "foo bar":
        var v41 int8 = 28
        _ = v41
    }
    if true {
        v3 = v3
    }
    var v42 float32 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
