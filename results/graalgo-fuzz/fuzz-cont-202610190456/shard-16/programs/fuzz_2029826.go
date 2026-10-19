package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint; F1 bool; F2 uint }
type GenS3 struct { F0 uint32 }
func genFn0(p0 bool, p1 string, p2 uint16) uint64 {
    switch 11 {
    case 2147483647:
        for fi5 := 0; fi5 < 19; fi5++ {
            switch (("hello" + (p1 + p1)) + "") {
            case "\n\t":
                p0 = p0
            default:
                v6 := []int{fi5, 4294967295, 127, -563}
                _ = v6
            }
            switch ("\n\t" + (p1 + ("" + p1))) {
            case "":
                var v7 uint = 4294967295
                _ = v7
            case "\n\t":
                v8 := []int{fi5, (fi5 ^ fi5), (fi5 / (((fi5 - fi5)) | 1))}
                _ = v8
            case "hello":
                p0 = true
            }
            _ = fi5
        }
    case -2147483648:
        v9 := []int{4294967295, 32767}
        _ = v9
    default:
        fmt.Println(p0, p1)
    }
    var v10 int32 = -574
    p2 -= (p2 % (((p2 % ((20) | 1))) | 1))
    if false {
        if true {
            var v11 uint32 = 2147483648
            p1 = ("foo bar" + p1)
            var v12 int16 = 0
            _ = v11
            _ = v12
        }
        v13 := []int{9223372036854775807, -2147483648, int(v10), 2147483647}
        _ = v13
    }
    v14 := []int{0, -278}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v14
    return 2
}

func main() {
    var v1 float32 = (0.0 / 1e10)
    var v2 int32 = -2147483648
    var v3 uint = 18446744073709551615
    v4 := []string{(("a" + "hello") + "\n\t"), "a", (("" + ("" + "\n\t")) + ("foo bar" + "a")), "\n\t"}
    if false {
        v1 = 0.1
        if true {
            fmt.Println(v1, v2, v1)
            var v15 bool = false
            _ = v15
        } else {
            if (0 <= 319) {
                v16 := GenS1{F0: 249, F1: false}
                _ = v16
            }
            for fi17 := 0; fi17 < 6; fi17++ {
                v2 = -1
                _ = fi17
            }
            var v18 uint32 = 67
            _ = v18
        }
        if (734 > 32767) {
            v19 := GenS1{F0: -1, F1: true}
            var v20 uint = v3
            v3 += v3
            _ = v19
            _ = v20
        }
    } else {
        v21 := GenS1{F0: -479, F1: true}
        if false {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch 9223372036854775807 {
            case 526:
                fmt.Println(v1, v3, v3)
            case -740:
                v4 = append(v4, "a")
            case 128:
                var v22 int64 = 42
                _ = v22
            default:
                switch "a" {
                case "a":
                    fmt.Println(v1)
                default:
                    switch "hello" {
                    case "":
                        v23 := GenS1{F0: -1, F1: (true && (true && true))}
                        _ = v23
                    case "foo bar":
                        var v24 float64 = (0.0 - 1.5)
                        _ = v24
                    case "a":
                        for fi25 := 0; fi25 < 18; fi25++ {
                            var v26 uint8 = 255
                            _ = fi25
                            _ = v26
                        }
                    }
                }
            }
        } else {
            if false {
                var v27 float32 = 0.1
                _ = v27
            } else {
                v28 := GenS1{F0: 9223372036854775807, F1: (v1 >= (1e10 + float32(0.1)))}
                v2 = 1
                _ = v28
            }
            for fi29 := 0; fi29 < 14; fi29++ {
                v30 := GenS2{F0: (v3 & (v3 / ((4294967295) | 1))), F1: false, F2: v3}
                _ = fi29
                _ = v30
            }
        }
        if len(v4) > 0 {
            v4[0] = ("a" + "a")
        }
        _ = v21
    }
    v1 += 1e10
    v3 += 21
    var v31 float32 = (v1 - v1)
    for fi32 := 0; fi32 < 18; fi32++ {
        fmt.Println(v2)
        var v33 int32 = v2
        if (9223372036854775807 < 464) {
            for fi34 := 0; fi34 < 12; fi34++ {
                switch (^(fi32)) {
                case 354:
                    v33 = v2
                case 9223372036854775807:
                    fmt.Println(fi34, fi34, v31)
                }
                fmt.Println(v33, v1)
                if false {
                    v3 = (v3 | (v3 % (((v3 & v3)) | 1)))
                }
                _ = fi34
            }
            var v35 float64 = 1.5
            v35 += float64(v31)
            _ = v35
        }
        _ = fi32
        _ = v33
    }
    v36 := GenS3{F0: 1}
    var v37 uint32 = 42
    var v38 int = 32768
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
