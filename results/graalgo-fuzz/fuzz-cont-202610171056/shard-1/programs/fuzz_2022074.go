package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 uint64; F2 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32 }
func genFn0(p0 int32) int8 {
    p0 = p0
    p0 = p0
    _ = p0
    return -1
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 42
    var v3 uint16 = (v2 + (v2 + v2))
    var v4 bool = (48 > 128)
    var v5 uint = 693
    v4 = false
    v6 := GenS1{F0: -9223372036854775808, F1: 0, F2: "\n\t"}
    var v7 uint8 = 255
    if ((v2 <= 2) || false) {
        if (v7 != (v7 | v7)) {
            fmt.Println(v2, v4)
            for fi8 := 0; fi8 < 12; fi8++ {
                if true {
                    v3 = v3
                    v3 += (v3 % (((^(v3))) | 1))
                    switch (fi8 % ((127) | 1)) {
                    case 317:
                        switch ("a" + "foo bar") {
                        case "":
                            v4 = v4
                        default:
                            _ = len(v1)
                        }
                    case 32767:
                        var v9 uint32 = 2147483648
                        _ = v9
                    case -80:
                        switch fi8 {
                        case -2147483648:
                            if (!((2 < 644) || (!v4))) {
                                for fi10 := 0; fi10 < 20; fi10++ {
                                    var v11 float32 = 3.14159
                                    v12 := GenS2{F0: v3}
                                    _ = fi10
                                    _ = v11
                                    _ = v12
                                }
                                var v13 uint16 = v2
                                _ = v13
                            }
                        }
                    }
                }
                _ = fi8
            }
        }
        var v14 float32 = -2.25
        _ = v14
    }
    switch "" {
    case "a":
        v15 := GenS1{F0: 128, F1: 1, F2: (("hello" + "") + ("foo bar" + "a"))}
        _ = v15
    case "":
        v2 = (v3 & 255)
    }
    fmt.Println(v5, v3, v2)
    var v16 int16 = -32768
    if v4 {
        v17 := GenS3{F0: 1, F1: 18446744073709551615, F2: -32}
        v2 = v2
        var v18 int = 2147483647
        _ = v17
        _ = v18
    }
    if (v4 || true) {
        fmt.Println(v4)
        var v19 uint8 = (v7 << ((v5) & 63))
        switch 4294967295 {
        case 32768:
            switch ("" + "\n\t") {
            case "hello":
                if v4 {
                    var v20 uint32 = 2147483648
                    _ = v20
                }
            default:
                v21 := v1[""]
                _ = v21
            }
        default:
            if v4 {
                fmt.Println(v3, v19)
            } else {
                fmt.Println(v5)
                var v22 bool = v4
                _ = v22
            }
        }
        _ = v19
    } else {
        if (0.0 != ((((0.1 * 0.1) - 1e10) / 1e10) / ((1e10 + (1.5 * 0.0)) - 0.0))) {
            switch "hello" {
            case "foo bar":
                for fi23 := 0; fi23 < 19; fi23++ {
                    var v24 int64 = 0
                    fmt.Println(fi23)
                    for fi25 := 0; fi25 < 18; fi25++ {
                        v26 := GenS3{F0: v24, F1: 509, F2: v16}
                        switch (-(fi23)) {
                        case -659:
                            var v27 int = fi23
                            v28 := GenS1{F0: v27, F1: -760, F2: ("\n\t" + "hello")}
                            _ = v27
                            _ = v28
                        case 2147483647:
                            var v29 bool = (!(true && true))
                            _ = v29
                        }
                        v30 := []int{fi23, -1, 159, (^(fi25))}
                        _ = fi25
                        _ = v26
                        _ = v30
                    }
                    _ = fi23
                    _ = v24
                }
            }
            var v31 bool = ((!(v4 || (1 <= -128))) || v4)
            _ = v31
        } else {
            fmt.Println(v4, v4)
            delete(v1, "foo bar")
            var v32 uint8 = (v7 | (^(v7)))
            _ = v32
        }
    }
    v33 := []int{1, 909, 1, 9223372036854775807}
    for fi34 := 0; fi34 < 10; fi34++ {
        v2 = v2
        var v35 string = (("foo bar" + "a") + "hello")
        _ = fi34
        _ = v35
    }
    for fi36 := 0; fi36 < 13; fi36++ {
        v37 := GenS4{F0: 0}
        if (true && ((1 != -974) && v4)) {
            v38 := GenS5{F0: 42}
            _ = v38
        }
        _ = fi36
        _ = v37
    }
    var v39 float64 = 0.1
    for fi40 := 0; fi40 < 2; fi40++ {
        fmt.Println(v39)
        _ = fi40
    }
    switch "hello" {
    case "a":
        var v41 bool = true
        _ = v41
    case "foo bar":
        v42 := GenS4{F0: 0}
        _ = v42
    }
    var v43 int32 = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v33
    _ = v39
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
}
