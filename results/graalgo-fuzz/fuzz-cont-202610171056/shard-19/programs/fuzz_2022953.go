package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 bool; F2 uint8 }
func genFn0(p0 uint64, p1 bool) int64 {
    v7 := GenS1{F0: (^(p0)), F1: "a"}
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 int64 = -129
        var v10 string = ("foo bar" + ("a" + ("" + "")))
        _ = fi8
        _ = v10
        _ = v9
    }
    if (p1 || (false && p1)) {
        v11 := []int{4294967295, 290, -1}
        for fi12 := 0; fi12 < 13; fi12++ {
            switch "a" {
            case "a":
                var v13 int32 = 2147483647
                _ = v13
            default:
                fmt.Println(fi12, p1, fi12)
            }
            _ = fi12
        }
        _ = v11
    } else {
        if (42 == 0) {
            for fi14 := 0; fi14 < 10; fi14++ {
                if (p1 && p1) {
                    for fi15 := 0; fi15 < 13; fi15++ {
                        var v16 float64 = (0.1 - 1.5)
                        v16 = v16
                        _ = fi15
                        _ = v16
                    }
                    var v17 int16 = -1
                    v18 := []int{fi14, -9223372036854775808, fi14, fi14}
                    _ = v17
                    _ = v18
                } else {
                    p0 = (p0 & p0)
                    fmt.Println(p1, p0)
                    switch fi14 {
                    case 32768:
                        switch (("foo bar" + "hello") + "hello") {
                        case "foo bar":
                            p1 = (p0 >= p0)
                        default:
                            for fi19 := 0; fi19 < 11; fi19++ {
                                v20 := []int{}
                                _ = fi19
                                _ = v20
                            }
                        }
                    }
                }
                var v21 bool = (!(1 != -4))
                p1 = false
                _ = fi14
                _ = v21
            }
        }
        fmt.Println(p0)
        var v22 string = "a"
        _ = v22
    }
    v23 := []int{32768, -1, 32767, -326}
    if len(v23) >= 2 {
        v23 = v23[1:len(v23)]
    }
    _ = p0
    _ = p1
    _ = v23
    _ = v7
    return 42
}

func main() {
    var v1 int32 = 1
    var v2 uint64 = 18446744073709551615
    var v3 uint = 939
    v4 := make(map[string]bool)
    var v5 int8 = 85
    var v6 float64 = 1e10
    for fi24 := 0; fi24 < 17; fi24++ {
        v25 := GenS1{F0: v2, F1: (("\n\t" + ("" + "hello")) + "")}
        switch "" {
        case "":
            if false {
                if false {
                    if (false && ("" > ("hello" + "hello"))) {
                        if true {
                            v26 := GenS1{F0: v2, F1: ("a" + "\n\t")}
                            _ = v26
                        }
                        var v27 int32 = v1
                        _ = v27
                    }
                    delete(v4, ("foo bar" + ((("\n\t" + "foo bar") + "a") + (("foo bar" + "a") + ("hello" + "foo bar")))))
                }
            }
        case "hello":
            v6 -= float64(-2.25)
        default:
            switch "\n\t" {
            case "a":
                var v28 string = ("foo bar" + "\n\t")
                _ = v28
            default:
                var v29 int16 = -421
                _ = v29
            }
        }
        _ = fi24
        _ = v25
    }
    for fi30 := 0; fi30 < 18; fi30++ {
        v1 += 42
        if (32767 == 78) {
            fmt.Println(fi30, v2, v3)
            v2 -= v2
        }
        _ = fi30
    }
    var v31 uint32 = 2
    fmt.Println(v6, v1, v1)
    var v32 uint8 = 2
    v6 -= v6
    if (!false) {
        for fi33 := 0; fi33 < 6; fi33++ {
            v34 := []int{}
            _ = fi33
            _ = v34
        }
        v4["hello"] = false
    }
    v35 := GenS2{F0: -9223372036854775808}
    switch 42 {
    case 127:
        v6 += v6
    case 32768:
        if false {
            v1 = v1
        } else {
            if (v31 >= v31) {
                fmt.Println(v32, v3, v5)
                v36 := GenS3{F0: (v1 + (v1 << (((v3 / ((uint(v31)) | 1))) & 63))), F1: true, F2: 1}
                if true {
                    v3 += (v3 << ((v3) & 63))
                } else {
                    var v37 uint16 = 1
                    _ = v37
                }
                _ = v36
            } else {
                v38 := []int{0, 1}
                _ = v38
            }
            fmt.Println(v6, v2)
            switch ("\n\t" + "\n\t") {
            case "hello":
                var v39 int = 127
                _ = v39
            case "a":
                var v40 int64 = 9223372036854775807
                _ = v40
            default:
                var v41 uint32 = 42
                _ = v41
            }
        }
    case -642:
        var v42 int32 = v1
        _ = v42
    }
    fmt.Println(v2)
    var v43 int16 = 258
    var v44 float64 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v35
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
