package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS3 struct { F0 int64; F1 bool; F2 int32 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 int64; F1 int64; F2 uint64 }
func genFn0(p0 bool, p1 int) uint16 {
    v4 := []int{}
    v5 := GenS1{F0: 127}
    switch ("\n\t" + "hello") {
    case "foo bar":
        for fi6 := 0; fi6 < 16; fi6++ {
            v7 := GenS2{F0: 1, F1: 472, F2: 2}
            var v8 string = "hello"
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    p0 = (p0 && true)
    if p0 {
        switch "\n\t" {
        case "":
            var v9 int8 = -114
            _ = v9
        case "foo bar":
            p1 += p1
        case "a":
            if p0 {
                fmt.Println(p1)
            } else {
                if p0 {
                    p1 = (p1 | (p1 >> ((4294967295) & 63)))
                }
            }
        }
        for fi10 := 0; fi10 < 14; fi10++ {
            p0 = (42 != -9223372036854775808)
            v11 := GenS3{F0: -1, F1: p0, F2: -650}
            for fi12 := 0; fi12 < 20; fi12++ {
                if p0 {
                    _ = len(v4)
                    p1 += 783
                    p0 = (p0 || (!(1 != 42)))
                }
                _ = fi12
            }
            _ = fi10
            _ = v11
        }
        v13 := GenS1{F0: 1}
        _ = v13
    } else {
        fmt.Println(p0)
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 269
}
func genFn1(p0 int32) int8 {
    switch 1 {
    case 2147483647:
        switch ("foo bar" + "foo bar") {
        case "a":
            var v14 uint32 = 65535
            _ = v14
        default:
            p0 = p0
        }
    default:
        p0 -= p0
    }
    p0 = (p0 >> ((2) & 63))
    p0 = p0
    v15 := []int{}
    _ = p0
    _ = v15
    return -128
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]int)
    var v3 int8 = -1
    var v16 int = 9223372036854775807
    v16 = (-(v16))
    if (!(0 < 2)) {
        switch 0 {
        case 128:
            v3 = (v3 / ((v3) | 1))
        }
        v3 += (-(v3))
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        switch (v16 % (((fi17 / (((fi17 >> ((65535) & 63))) | 1))) | 1)) {
        case 32767:
            v3 = (-(v3))
        default:
            for fi18 := 0; fi18 < 11; fi18++ {
                var v19 uint16 = 78
                _ = fi18
                _ = v19
            }
        }
        _ = fi17
    }
    if false {
        if true {
            switch "" {
            case "hello":
                if ((1e10 / 1.5) == 1e10) {
                    fmt.Println(v3)
                    delete(v1, "hello")
                } else {
                    v20 := GenS1{F0: 2}
                    for fi21 := 0; fi21 < 15; fi21++ {
                        var v22 int64 = -255
                        _ = fi21
                        _ = v22
                    }
                    if (!(!((!(!true)) && true))) {
                        switch (^(v16)) {
                        case -530:
                            switch ((("a" + "") + "foo bar") + ("hello" + (((("a" + "") + "\n\t") + "a") + "hello"))) {
                            case "hello":
                                var v23 uint8 = 228
                                _ = v23
                            }
                        case -9223372036854775808:
                            if false {
                                var v24 uint32 = 385
                                _ = v24
                            }
                        case -736:
                            v25 := []int{(-(v16)), v16, -6, (v16 >> ((42) & 63))}
                            _ = v25
                        }
                    }
                    _ = v20
                }
            case "foo bar":
                v16 -= 127
            }
            v26 := []int{}
            _ = v26
        }
    } else {
        if true {
            v3 -= genFn1(1)
            var v27 uint64 = 18446744073709551615
            _ = v27
        }
    }
    delete(v2, 475)
    if (4294967295 < 0) {
        var v28 uint = 65535
        if (("a" + ("" + "hello")) < "") {
            v29 := []int{(v16 % (((v16 + -1)) | 1)), v16, v16, (v16 * 32767)}
            if false {
                for fi30 := 0; fi30 < 20; fi30++ {
                    if (933 != 42) {
                        for fi31 := 0; fi31 < 13; fi31++ {
                            var v32 int32 = -24
                            _ = fi31
                            _ = v32
                        }
                    }
                    _ = fi30
                }
            } else {
                v33 := GenS4{F0: 65535, F1: "hello"}
                _ = v33
            }
            _ = v29
        }
        v34 := GenS1{F0: 107}
        _ = v28
        _ = v34
    } else {
        if ((!(!true)) && (!true)) {
            fmt.Println(v16, v3)
            var v35 uint32 = 2147483647
            v16 += v16
            _ = v35
        }
        delete(v1, "\n\t")
        v16 += (v16 & v16)
    }
    var v36 int32 = 953
    var v37 uint32 = 841
    fmt.Println(v16, v37, v3)
    for fi38 := 0; fi38 < 10; fi38++ {
        var v39 int = (fi38 & (fi38 & -1))
        fmt.Println(v37)
        _ = fi38
        _ = v39
    }
    var v40 int32 = -2147483648
    v41 := GenS5{F0: 42, F1: 1, F2: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v40
    _ = v41
}
