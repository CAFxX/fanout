package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint }
func genFn0(p0 int32, p1 uint) bool {
    var v6 uint32 = 35
    var v7 int8 = 1
    switch 32768 {
    case 196:
        v8 := GenS1{F0: p1, F1: 1, F2: -1}
        _ = v8
    case 32767:
        p1 = 42
    }
    if ((128 > 160) || false) {
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 uint16 = 32767
            p0 -= p0
            _ = fi9
            _ = v10
        }
        v11 := []int{}
        _ = v11
    } else {
        v12 := []int{}
        switch ("hello" + "\n\t") {
        case "a":
            fmt.Println(p0, v6)
        case "foo bar":
            var v13 int64 = -329
            _ = v13
        }
        _ = v12
    }
    var v14 uint = (p1 & 484)
    _ = p0
    _ = p1
    _ = v14
    _ = v6
    _ = v7
    return (0 == -1)
}
func genFn1() uint8 {
    if (!false) {
        v15 := GenS2{F0: 757}
        _ = v15
    }
    if (!true) {
        if true {
            var v16 bool = (true && (1.5 >= 0.0))
            fmt.Println(v16)
            switch "hello" {
            case "hello":
                var v17 int16 = -582
                _ = v17
            case "":
                v18 := GenS2{F0: 0}
                _ = v18
            default:
                var v19 int32 = -2147483648
                _ = v19
            }
            _ = v16
        } else {
            var v20 float64 = 0.1
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        if false {
            var v22 int8 = -73
            _ = v22
        }
        _ = fi21
    }
    return 42
}

func main() {
    v1 := []string{"", "foo bar"}
    v2 := make(map[uint32]int)
    var v3 int32 = 280
    v4 := []string{}
    var v5 float32 = 1.5
    v5 += float32(float64((v5 * float32((1.5 + -2.25)))))
    v3 = v3
    var v23 bool = true
    var v24 string = (("foo bar" + ("\n\t" + ("" + ""))) + "")
    for fi25 := 0; fi25 < 20; fi25++ {
        v5 -= float32(0.1)
        var v26 bool = true
        var v27 float64 = 0.0
        _ = fi25
        _ = v26
        _ = v27
    }
    if true {
        if (248 != 0) {
            _ = len(v2)
            for fi28 := 0; fi28 < 6; fi28++ {
                if (!(!false)) {
                    fmt.Println(v24)
                    fmt.Println(v24, v23)
                }
                if true {
                    _ = len(v1)
                    v29 := GenS1{F0: 2, F1: 127, F2: 42}
                    _ = v29
                } else {
                    if true {
                        v24 = (("hello" + (("\n\t" + v24) + v24)) + (v24 + "hello"))
                        _ = len(v2)
                    }
                    fmt.Println(v24, v3)
                }
                if (753 != 1000) {
                    var v30 int = 1
                    v23 = (212 < 568)
                    _ = v30
                } else {
                    v24 = v24
                }
                _ = fi28
            }
        }
    } else {
        switch (v24 + v24) {
        case "\n\t":
            if ((true && (!(((6 > genFn1()) && v23) && true))) || v23) {
                v31 := GenS2{F0: 65535}
                v3 = (v3 % ((v3) | 1))
                _ = v31
            } else {
                for fi32 := 0; fi32 < 20; fi32++ {
                    var v33 float64 = 1e10
                    if true {
                        _ = len(v1)
                        v24 = v24
                    } else {
                        v34 := GenS2{F0: 32767}
                        fmt.Println(v5, v24, v24)
                        _ = v34
                    }
                    _ = fi32
                    _ = v33
                }
            }
        case "":
            if true {
                for fi35 := 0; fi35 < 16; fi35++ {
                    if v23 {
                        fmt.Println(v5, fi35, v5)
                        if false {
                            v3 -= -357
                            var v36 uint16 = 2
                            _ = v36
                        } else {
                            var v37 uint8 = 127
                            _ = v37
                        }
                        var v38 float32 = -2.25
                        _ = v38
                    }
                    v39 := GenS2{F0: 255}
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi35
                    _ = v39
                }
                if v23 {
                    var v40 uint8 = genFn1()
                    v5 += 3.14159
                    _ = v40
                }
            }
        case "hello":
            fmt.Println(v24)
        }
        v41 := v2[483]
        _ = v41
    }
    delete(v2, 2)
    if false {
        for fi42 := 0; fi42 < 10; fi42++ {
            switch (("" + (v24 + ("foo bar" + v24))) + "") {
            case "\n\t":
                v43 := GenS3{F0: fi42, F1: 896}
                _ = v43
            default:
                _ = len(v2)
            }
            _ = fi42
        }
    } else {
        for fi44 := 0; fi44 < 17; fi44++ {
            switch (((v24 + v24) + (v24 + v24)) + v24) {
            case "hello":
                for fi45 := 0; fi45 < 14; fi45++ {
                    v46 := GenS2{F0: 374}
                    if v23 {
                        var v47 int8 = -128
                        v5 = 1.5
                        _ = v47
                    } else {
                        fmt.Println(v24)
                        v48 := GenS1{F0: 1, F1: -128, F2: -9223372036854775808}
                        v5 -= v5
                        _ = v48
                    }
                    _ = fi45
                    _ = v46
                }
            case "a":
                var v49 int64 = 0
                _ = v49
            case "":
                for fi50 := 0; fi50 < 2; fi50++ {
                    v24 = v24
                    _ = len(v1)
                    _ = fi50
                }
            }
            for fi51 := 0; fi51 < 8; fi51++ {
                v52 := GenS2{F0: 371}
                _ = fi51
                _ = v52
            }
            switch "" {
            case "":
                var v53 uint32 = 883
                _ = v53
            case "a":
                if len(v1) > 0 {
                    v1[0] = v24
                }
            }
            _ = fi44
        }
        v23 = v23
    }
    v54 := GenS3{F0: 4294967295, F1: 2}
    v23 = v23
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v54
}
