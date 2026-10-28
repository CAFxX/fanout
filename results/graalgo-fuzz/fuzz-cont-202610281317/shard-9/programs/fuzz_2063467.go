package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint8; F2 uint16 }
type GenS3 struct { F0 uint32; F1 bool; F2 uint16 }
func genFn0(p0 int8, p1 uint8, p2 uint) int32 {
    var v4 int = -9223372036854775808
    p0 -= 100
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 2147483647
}
func genFn1(p0 uint) int {
    for fi5 := 0; fi5 < 3; fi5++ {
        for fi6 := 0; fi6 < 10; fi6++ {
            var v7 string = "foo bar"
            var v8 float32 = ((3.14159 / float32((-2.25 * 1e10))) + (1.5 - ((1e10 - 0.0) * (float32(3.14159) / 1.0))))
            _ = fi6
            _ = v7
            _ = v8
        }
        switch "a" {
        case "hello":
            for fi9 := 0; fi9 < 2; fi9++ {
                var v10 string = "a"
                if (!true) {
                    if false {
                        if true {
                            v11 := GenS1{F0: 2147483648}
                            _ = v11
                        }
                        p0 -= 18446744073709551615
                    } else {
                        var v12 int32 = 1
                        _ = v12
                    }
                } else {
                    v13 := []int{42, -1, fi5, fi5}
                    _ = v13
                }
                _ = fi9
                _ = v10
            }
        case "\n\t":
            var v14 uint = (p0 - p0)
            _ = v14
        default:
            p0 += (p0 % ((p0) | 1))
        }
        for fi15 := 0; fi15 < 2; fi15++ {
            v16 := GenS1{F0: 801}
            _ = fi15
            _ = v16
        }
        _ = fi5
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        if (!(((2147483647 >= 42) && ((((365 < 667) && true) || true) || (1 != 42))) && false)) {
            v18 := []int{fi17, fi17, (-(fi17))}
            _ = v18
        } else {
            fmt.Println(p0, p0)
            p0 += p0
        }
        p0 += p0
        _ = fi17
    }
    var v19 uint16 = 2
    v20 := []int{59, -9223372036854775808, -1, 32767}
    for fi21 := 0; fi21 < 9; fi21++ {
        var v22 int8 = 1
        v23 := GenS2{F0: genFn0((v22 * v22), 233, p0), F1: 238, F2: v19}
        _ = fi21
        _ = v22
        _ = v23
    }
    _ = p0
    _ = v19
    _ = v20
    return -1
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{"hello"}
    v3 := make(map[string]bool)
    if (false || false) {
        v24 := GenS3{F0: 2, F1: ((float32(-2.25) - 3.14159) >= float32((1e10 / 3.14159))), F2: 712}
        _ = v24
    }
    delete(v1, "hello")
    for fi25 := 0; fi25 < 1; fi25++ {
        var v26 bool = (!(808 <= 65535))
        var v27 string = ""
        _ = len(v2)
        _ = fi25
        _ = v26
        _ = v27
    }
    switch 127 {
    case 2147483647:
        switch -859 {
        case 42:
            v28 := GenS2{F0: 42, F1: 0, F2: 42}
            _ = v28
        case -1:
            if ((286 > 42) || true) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                delete(v1, "foo bar")
                switch "\n\t" {
                case "hello":
                    var v29 uint = 262
                    _ = v29
                case "\n\t":
                    v30 := v3[("" + "")]
                    _ = v30
                }
            }
        default:
            var v31 float64 = float64(1e10)
            _ = v31
        }
    case 128:
        var v32 string = ("foo bar" + ("foo bar" + "foo bar"))
        _ = v32
    }
    v33 := GenS2{F0: 0, F1: 255, F2: 2}
    for fi34 := 0; fi34 < 13; fi34++ {
        var v35 string = "a"
        for fi36 := 0; fi36 < 3; fi36++ {
            v37 := v3[v35]
            _ = fi36
            _ = v37
        }
        _ = fi34
        _ = v35
    }
    if false {
        var v38 int16 = -32768
        v38 -= 42
        _ = v38
    }
    var v39 float32 = float32(3.14159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v39
}
