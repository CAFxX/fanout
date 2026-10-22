package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS2 struct { F0 string; F1 uint32 }
func genFn0() uint {
    v5 := []int{-9223372036854775808, 32768}
    var v6 bool = (481 <= 2)
    if true {
        switch -9223372036854775808 {
        case 643:
            var v7 int64 = 1
            _ = v7
        default:
            switch ("" + "hello") {
            case "\n\t":
                for fi8 := 0; fi8 < 1; fi8++ {
                    fmt.Println(v6)
                    _ = fi8
                }
            case "hello":
                if (((3.14159 * 1e10) == (1e10 - 0.1)) && v6) {
                    if (42 <= 1) {
                        fmt.Println(v6)
                        fmt.Println(v6)
                        var v9 uint8 = 128
                        _ = v9
                    } else {
                        for fi10 := 0; fi10 < 20; fi10++ {
                            if true {
                                v6 = (1 <= 255)
                                v6 = true
                            }
                            var v11 int16 = 1000
                            _ = fi10
                            _ = v11
                        }
                    }
                    var v12 int16 = 84
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v12
                }
            case "a":
                fmt.Println(v6)
            }
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 bool) uint32 {
    var v13 uint8 = 2
    switch 0 {
    case 127:
        v14 := GenS1{F0: v13, F1: 4294967295, F2: 42}
        _ = v14
    case 0:
        v15 := GenS2{F0: "foo bar", F1: 0}
        _ = v15
    case 128:
        if p0 {
            if (3.14159 >= 0.1) {
                var v16 uint64 = 18446744073709551615
                _ = v16
            }
            switch "a" {
            case "a":
                for fi17 := 0; fi17 < 1; fi17++ {
                    v13 = (v13 % (((v13 | v13)) | 1))
                    var v18 bool = (2147483648 <= 4294967295)
                    v18 = (("a" + ((("\n\t" + "foo bar") + ("hello" + "a")) + "\n\t")) > ("a" + ""))
                    _ = fi17
                    _ = v18
                }
            }
            if p0 {
                for fi19 := 0; fi19 < 4; fi19++ {
                    p0 = false
                    switch (fi19 + (fi19 % ((-665) | 1))) {
                    case 9223372036854775807:
                        p0 = p0
                    case -1:
                        fmt.Println(p0, p0, v13)
                    }
                    _ = fi19
                }
            }
        }
    }
    _ = p0
    _ = v13
    return 4294967295
}

func main() {
    var v1 int8 = 1
    v2 := []string{}
    var v3 bool = false
    var v4 float64 = 3.14159
    var v20 int8 = 42
    if false {
        v20 -= v1
    }
    v21 := GenS1{F0: 128, F1: genFn1(v3), F2: genFn1((!v3))}
    for fi22 := 0; fi22 < 9; fi22++ {
        for fi23 := 0; fi23 < 8; fi23++ {
            v24 := GenS2{F0: "foo bar", F1: 2147483647}
            var v25 uint32 = 0
            var v26 uint32 = (v25 | (^(v25)))
            fmt.Println(v4, v20, v1)
            switch "hello" {
            case "":
                if (-2147483648 < -692) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v27 uint16 = 32767
                    _ = v27
                } else {
                    fmt.Println(v3)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v25 = v25
                    v1 += -63
                }
            }
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
        v4 -= float64(3.14159)
        _ = fi22
    }
    var v28 string = ""
    switch "a" {
    case "hello":
        fmt.Println(v3, v28, v3)
    }
    if false {
        if len(v2) > 0 {
            v2[0] = v28
        }
        var v29 int = 2147483647
        _ = v29
    } else {
        var v30 string = ("\n\t" + ((v28 + v28) + v28))
        v1 = v20
        var v31 int8 = (-(v20))
        _ = v30
        _ = v31
    }
    switch 128 {
    case 42:
        var v32 uint64 = 0
        _ = v32
    case 32767:
        if false {
            var v33 float32 = 1e10
            if (v20 >= v1) {
                var v34 uint32 = 4294967295
                _ = v34
            } else {
                if (!v3) {
                    if len(v2) > 0 {
                        v2[0] = v28
                    }
                    v20 = 100
                } else {
                    if len(v2) > 0 {
                        v2[0] = ("hello" + (v28 + ""))
                    }
                }
                v4 -= float64(-2.25)
                var v35 int64 = 42
                _ = v35
            }
            _ = v33
        } else {
            fmt.Println(v4, v20, v1)
        }
    case -85:
        v20 += (v20 + v20)
    }
    var v36 uint32 = 2147483648
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v37 int8 = (v20 >> ((4294967295) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
