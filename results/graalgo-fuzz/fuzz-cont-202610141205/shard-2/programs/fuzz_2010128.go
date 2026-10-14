package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int; F1 uint8; F2 int32 }
func genFn0() string {
    switch ("a" + ("hello" + "foo bar")) {
    case "":
        var v6 uint32 = 42
        _ = v6
    default:
        var v7 uint16 = 65535
        _ = v7
    }
    v8 := GenS1{F0: 18446744073709551615, F1: -1}
    var v9 int32 = 42
    _ = v8
    _ = v9
    return "hello"
}
func genFn1(p0 uint8, p1 uint64, p2 bool) int32 {
    fmt.Println(p0, p0, p2)
    fmt.Println(p0)
    p1 = 4294967295
    switch "" {
    case "foo bar":
        p1 = p1
    case "\n\t":
        if (false && (!true)) {
            var v10 int64 = 9223372036854775807
            _ = v10
        }
    default:
        for fi11 := 0; fi11 < 8; fi11++ {
            switch ("\n\t" + "hello") {
            case "foo bar":
                p1 = p1
            case "\n\t":
                p1 -= 18446744073709551615
            case "a":
                if true {
                    var v12 uint16 = 147
                    fmt.Println(p2, v12, v12)
                    var v13 int16 = -643
                    _ = v12
                    _ = v13
                }
            }
            v14 := []int{9223372036854775807}
            _ = fi11
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return -2147483648
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 166
    var v3 uint64 = 4294967295
    var v4 uint = uint(v2)
    var v5 bool = (!(false && ((0.0 > (1e10 - float32(0.1))) && false)))
    var v15 uint8 = v2
    v3 += (v3 / ((v3) | 1))
    if ((!(65535 == 31)) && v5) {
        for fi16 := 0; fi16 < 5; fi16++ {
            for fi17 := 0; fi17 < 3; fi17++ {
                v3 -= uint64(v4)
                v18 := []int{}
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
    } else {
        v19 := GenS2{F0: 75, F1: (v2 << (((v4 % (((^(v4))) | 1))) & 63)), F2: 816}
        _ = v19
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 int16 = 862
        v5 = (v5 && (1 > v4))
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        for fi23 := 0; fi23 < 5; fi23++ {
            v24 := []int{(^(fi22))}
            if (true || v5) {
                var v25 float64 = float64((1e10 / -2.25))
                switch 32768 {
                case 128:
                    var v26 string = "a"
                    _ = v26
                }
                v4 = (^(v4))
                _ = v25
            } else {
                var v27 int = fi23
                var v28 int32 = genFn1((v15 | v15), (^(v3)), false)
                v28 += int32(fi22)
                _ = v27
                _ = v28
            }
            fmt.Println(v5, fi23, fi23)
            _ = fi23
            _ = v24
        }
        if (genFn1((^(v2)), v3, (255 >= 326)) > 884) {
            v29 := []int{127, fi22}
            var v30 float64 = float64(0.0)
            var v31 int64 = 42
            _ = v29
            _ = v30
            _ = v31
        }
        var v32 float32 = (float32(float64((((0.1 * 1.5) + (0.1 / (((-2.25 - 1e10) + -2.25) * float32((float64(((-2.25 + 1e10) * 3.14159)) - ((((1.5 * 1e10) / 1.0) * 3.14159) / ((float64(float32(-2.25)) * 0.1) * ((3.14159 + 3.14159) + 3.14159)))))))) * 0.0))) - (0.1 + ((float32((1e10 / 1.5)) / 0.1) - ((float32(-2.25) - float32((float64(1.5) - (1e10 / 1.5)))) * float32(((3.14159 + (0.1 - (0.1 * ((-2.25 / 1.0) / (1.5 / -2.25))))) + 1.5))))))
        _ = fi22
        _ = v32
    }
    v33 := []int{0}
    v5 = v5
    for fi34 := 0; fi34 < 8; fi34++ {
        switch "foo bar" {
        case "hello":
            v3 = (v3 + 193)
        case "\n\t":
            fmt.Println(v4, v5)
        case "":
            if true {
                v5 = (14 <= 15)
            } else {
                var v35 uint16 = 2
                _ = v35
            }
        default:
            fmt.Println(v15)
        }
        var v36 string = "\n\t"
        v1[(v36 + v36)] = true
        _ = fi34
        _ = v36
    }
    switch "" {
    case "hello":
        var v37 uint64 = (v3 + 2)
        _ = v37
    case "":
        for fi38 := 0; fi38 < 18; fi38++ {
            v2 -= v2
            _ = fi38
        }
    case "\n\t":
        v5 = (756 >= 2)
    default:
        for fi39 := 0; fi39 < 15; fi39++ {
            v5 = (!((v2 % (((v15 << (((^(v4))) & 63))) | 1)) <= (v2 % (((v15 - 127)) | 1))))
            v1["hello"] = false
            v2 -= (v15 >> ((v4) & 63))
            _ = fi39
        }
    }
    v40 := GenS1{F0: v4, F1: 1}
    v15 -= 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v40
    _ = v5
}
