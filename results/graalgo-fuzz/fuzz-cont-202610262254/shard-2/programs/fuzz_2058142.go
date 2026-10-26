package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 string; F1 uint8 }
func genFn0(p0 uint32, p1 bool) int8 {
    p1 = false
    if false {
        p0 = p0
        var v6 bool = (!(!p1))
        _ = v6
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        if ((true && p1) && (128 > 0)) {
            v8 := []int{fi7, 127}
            _ = v8
        }
        var v9 bool = p1
        _ = fi7
        _ = v9
    }
    _ = p0
    _ = p1
    return 38
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = (0.0 + 3.14159)
    v3 := make(map[uint32]bool)
    v4 := []uint32{688}
    v5 := []int{27, 32767}
    switch "a" {
    case "foo bar":
        v2 = v2
    case "a":
        fmt.Println(v2)
    }
    v2 += 0.0
    v10 := GenS1{F0: 42}
    for fi11 := 0; fi11 < 6; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            v2 = v2
            v13 := GenS1{F0: 455}
            switch (("\n\t" + "") + ("\n\t" + (("\n\t" + "a") + (("a" + ("" + ("" + "hello"))) + "hello")))) {
            case "hello":
                v14 := GenS2{F0: -540, F1: 9223372036854775807}
                _ = v14
            case "a":
                v2 += v2
            case "\n\t":
                delete(v3, 431)
            default:
                if (!(226 > 38)) {
                    v15 := GenS1{F0: 42}
                    v2 += v2
                    fmt.Println(v2)
                    _ = v15
                } else {
                    switch (fi11 / ((fi11) | 1)) {
                    case -2147483648:
                        var v16 int8 = genFn0(42, false)
                        switch 32767 {
                        case 2147483647:
                            var v17 float32 = (1.5 / ((v2 - 3.14159) * (((-2.25 * v2) / v2) * float32(float64((1.5 - (float32(3.14159) - (((((1e10 / ((v2 / (v2 * (v2 / ((0.0 * 1e10) * v2)))) + 1e10)) - (((3.14159 / v2) - 1e10) / 1e10)) / v2) / 0.1) * (float32(3.14159) - (v2 - v2))))))))))
                            _ = v17
                        case -760:
                            v16 = 0
                        case 128:
                            v18 := GenS2{F0: 1, F1: 456}
                            _ = v18
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if (255 != 18446744073709551615) {
                            v3[42] = true
                            v2 += v2
                        }
                        v19 := GenS3{F0: "a", F1: 22}
                        _ = v16
                        _ = v19
                    }
                    var v20 int = (fi12 & fi12)
                    _ = len(v1)
                    _ = v20
                }
            }
            _ = fi12
            _ = v13
        }
        switch "foo bar" {
        case "":
            var v21 string = "a"
            _ = v21
        case "hello":
            v2 = -2.25
        case "\n\t":
            if ("hello" != "a") {
                var v22 int32 = -1
                var v23 int64 = 0
                if (((^(fi11)) > (fi11 | 42)) && (("\n\t" + "\n\t") == "foo bar")) {
                    var v24 uint = 914
                    var v25 uint16 = 65535
                    if len(v5) > 0 {
                        v5[0] = fi11
                    }
                    _ = v24
                    _ = v25
                } else {
                    var v26 uint32 = 42
                    _ = v26
                }
                _ = v22
                _ = v23
            }
        }
        _ = fi11
    }
    if false {
        for fi27 := 0; fi27 < 7; fi27++ {
            v2 -= (((0.0 * v2) + (v2 + v2)) * v2)
            var v28 int32 = -2147483648
            _ = fi27
            _ = v28
        }
        switch "a" {
        case "a":
            v2 += v2
        case "":
            v2 = v2
        case "hello":
            if false {
                v5 = append(v5, 2147483647)
            }
        }
    }
    v2 = ((v2 * float32(1e10)) - (v2 - (((v2 + v2) + (float32(0.0) * 1e10)) + v2)))
    v29 := GenS2{F0: 2147483647, F1: 9223372036854775807}
    v30 := GenS2{F0: -724, F1: 346}
    switch ("hello" + (((("hello" + ("foo bar" + ((("" + "hello") + "a") + ("a" + "")))) + "") + "hello") + ("\n\t" + ("" + "")))) {
    case "hello":
        for fi31 := 0; fi31 < 19; fi31++ {
            var v32 uint = 18446744073709551615
            _ = fi31
            _ = v32
        }
    }
    var v33 string = "a"
    v34 := GenS2{F0: -1, F1: 964}
    var v35 int8 = 14
    fmt.Println(v35)
    for fi36 := 0; fi36 < 11; fi36++ {
        fmt.Println(v33)
        v35 -= v35
        var v37 int64 = 486
        _ = fi36
        _ = v37
    }
    var v38 uint = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v35
    _ = v38
    _ = v4
    _ = v5
}
