package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint; F2 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 uint32 }
func genFn0(p0 bool) int {
    for fi5 := 0; fi5 < 8; fi5++ {
        p0 = false
        var v6 int8 = 100
        _ = fi5
        _ = v6
    }
    if true {
        var v7 int32 = 42
        var v8 uint = 2
        _ = v7
        _ = v8
    }
    fmt.Println(p0)
    p0 = (p0 && p0)
    v9 := []int{9223372036854775807, -459, -9223372036854775808, 0}
    _ = p0
    _ = v9
    return 2147483647
}

func main() {
    var v1 bool = (2147483647 != -1)
    v2 := []int{4294967295, 4294967295, 2147483647}
    v3 := []int{-463}
    v4 := make(map[string]bool)
    switch 32768 {
    case 32768:
        if (!v1) {
            v4[(("a" + ("" + "\n\t")) + "a")] = (127 != genFn0(v1))
        } else {
            v1 = ((float32(-2.25) != ((-2.25 + 0.0) + 3.14159)) || v1)
            switch 9223372036854775807 {
            case 2147483647:
                var v10 uint8 = 42
                _ = v10
            default:
                switch genFn0((v1 || false)) {
                case 849:
                    v2 = append(v2, genFn0((25 != 128)))
                case 9223372036854775807:
                    delete(v4, ("hello" + ("\n\t" + "\n\t")))
                case 1:
                    v11 := v4[(((("hello" + "a") + "") + "a") + "hello")]
                    _ = v11
                default:
                    var v12 string = "a"
                    _ = v12
                }
            }
        }
    case 127:
        _ = len(v3)
    default:
        var v13 int64 = 1
        _ = v13
    }
    var v14 int8 = -128
    v1 = ((-2147483648 < 0) || v1)
    _ = len(v4)
    fmt.Println(v14)
    switch -9223372036854775808 {
    case 2147483647:
        switch "foo bar" {
        case "":
            var v15 int8 = -60
            _ = v15
        case "a":
            for fi16 := 0; fi16 < 6; fi16++ {
                fmt.Println(fi16, v1, v1)
                fmt.Println(fi16, fi16, fi16)
                _ = fi16
            }
        }
    case 511:
        _ = len(v4)
    case 934:
        v1 = (65535 > 850)
    }
    switch -2147483648 {
    case 314:
        if v1 {
            var v17 float64 = 1e10
            v18 := GenS1{F0: 255}
            var v19 float64 = v17
            _ = v17
            _ = v18
            _ = v19
        } else {
            v20 := GenS2{F0: 1, F1: 269, F2: 0}
            for fi21 := 0; fi21 < 15; fi21++ {
                switch (fi21 & 4294967295) {
                case 4294967295:
                    var v22 float64 = 0.1
                    _ = v22
                case 0:
                    v14 += (v14 * -11)
                case -9223372036854775808:
                    v23 := GenS1{F0: 1}
                    _ = v23
                }
                var v24 bool = (!(!v1))
                for fi25 := 0; fi25 < 4; fi25++ {
                    v24 = true
                    _ = fi25
                }
                _ = fi21
                _ = v24
            }
            _ = v20
        }
    case -2147483648:
        switch genFn0(v1) {
        case -9223372036854775808:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 2147483647:
            for fi26 := 0; fi26 < 19; fi26++ {
                switch ("hello" + (("" + ("" + ("foo bar" + ""))) + "\n\t")) {
                case "foo bar":
                    fmt.Println(v14)
                case "":
                    var v27 uint32 = 2147483647
                    _ = v27
                case "hello":
                    if ((!v1) && (v1 || (!false))) {
                        var v28 uint8 = 189
                        var v29 int16 = 1
                        var v30 float64 = 3.14159
                        _ = v28
                        _ = v29
                        _ = v30
                    } else {
                        var v31 float32 = 1.5
                        var v32 uint8 = 42
                        v4["hello"] = true
                        _ = v31
                        _ = v32
                    }
                default:
                    if false {
                        v1 = (v1 || ((!true) || v1))
                    } else {
                        v14 -= v14
                        for fi33 := 0; fi33 < 9; fi33++ {
                            var v34 uint64 = 18446744073709551615
                            _ = fi33
                            _ = v34
                        }
                        var v35 int16 = int16(v14)
                        _ = v35
                    }
                }
                switch (fi26 - fi26) {
                case 502:
                    fmt.Println(fi26)
                default:
                    switch (fi26 & fi26) {
                    case -2147483648:
                        var v36 uint8 = 49
                        _ = v36
                    case 42:
                        var v37 int16 = 42
                        _ = v37
                    default:
                        v38 := GenS3{F0: genFn0(v1)}
                        _ = v38
                    }
                }
                _ = fi26
            }
        default:
            for fi39 := 0; fi39 < 18; fi39++ {
                if v1 {
                    v4[(((("\n\t" + "foo bar") + "a") + "\n\t") + "a")] = v1
                    v40 := GenS4{F0: (1e10 >= 0.0), F1: 0}
                    v1 = true
                    _ = v40
                }
                _ = fi39
            }
        }
    default:
        v1 = v1
    }
    if v1 {
        fmt.Println(v1)
    }
    if (4294967295 != 1) {
        var v41 string = "foo bar"
        _ = v41
    } else {
        v14 = v14
        var v42 float64 = 0.0
        v42 = v42
        _ = v42
    }
    fmt.Println(v1, v14)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
