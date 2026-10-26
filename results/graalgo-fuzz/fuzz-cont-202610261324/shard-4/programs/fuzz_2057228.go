package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 int64; F2 uint32 }
func genFn0() int8 {
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 string = ((("hello" + "") + "a") + ("" + ("foo bar" + (("hello" + "hello") + "\n\t"))))
        v9 := []int{(fi7 - 4294967295), fi7, 42}
        v8 = "hello"
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 1, F1: 954}
    switch "a" {
    case "foo bar":
        var v11 int8 = 127
        _ = v11
    }
    switch -1 {
    case -9223372036854775808:
        var v12 bool = false
        _ = v12
    case 127:
        v13 := []int{-309, 192}
        _ = v13
    }
    _ = v10
    return 122
}
func genFn1(p0 int, p1 uint16) int {
    v14 := GenS1{F0: -342, F1: 122}
    var v15 uint = 0
    v16 := []int{}
    p0 -= -472
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v16
    return p0
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]uint32)
    v3 := make(map[int]string)
    v4 := make(map[string]uint32)
    var v5 string = (("\n\t" + "hello") + (("\n\t" + ("hello" + "hello")) + "foo bar"))
    var v6 string = v5
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 uint = 973
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        for fi20 := 0; fi20 < 16; fi20++ {
            switch (v5 + v5) {
            case "a":
                v21 := GenS1{F0: -274, F1: 480}
                _ = v21
            case "\n\t":
                v5 = v6
                v6 = "hello"
                if false {
                    var v22 string = (v5 + v5)
                    if (65535 > 42) {
                        v23 := GenS1{F0: 0, F1: 2}
                        v24 := GenS1{F0: -763, F1: 4294967295}
                        if (!(("foo bar" + "foo bar") != ((v5 + "hello") + "foo bar"))) {
                            var v25 uint32 = 0
                            var v26 int16 = 0
                            _ = v25
                            _ = v26
                        }
                        _ = v23
                        _ = v24
                    } else {
                        v5 = v22
                        fmt.Println(fi19, v5, fi20)
                    }
                    v22 = ""
                    v27 := []int{}
                    _ = v22
                    _ = v27
                } else {
                    v28 := v4[v6]
                    var v29 float32 = (0.0 / 0.1)
                    fmt.Println(fi20)
                    _ = v28
                    _ = v29
                }
            default:
                v6 = v5
                fmt.Println(fi20, v6)
                delete(v3, (fi20 + (fi19 % (((^(fi19))) | 1))))
            }
            _ = fi20
        }
        _ = fi19
    }
    for fi30 := 0; fi30 < 1; fi30++ {
        if true {
            switch fi30 {
            case 376:
                fmt.Println(v6, v6)
            case 42:
                var v31 uint8 = 0
                _ = v31
            }
        }
        _ = fi30
    }
    v6 = "a"
    for fi32 := 0; fi32 < 11; fi32++ {
        v33 := []int{}
        _ = fi32
        _ = v33
    }
    for fi34 := 0; fi34 < 8; fi34++ {
        v6 = v5
        if (true && true) {
            var v35 string = ("\n\t" + v6)
            for fi36 := 0; fi36 < 19; fi36++ {
                var v37 float64 = 0.0
                v38 := []int{(fi34 << ((235) & 63))}
                _ = fi36
                _ = v37
                _ = v38
            }
            _ = v35
        } else {
            switch "" {
            case "foo bar":
                v39 := v2[fi34]
                _ = v39
            case "\n\t":
                v40 := []int{2147483647, (fi34 << ((255) & 63)), fi34}
                _ = v40
            case "":
                if (("" <= v5) || (2147483647 >= -966)) {
                    if (!(!(float32(-2.25) == 1e10))) {
                        v41 := GenS1{F0: -1, F1: 2}
                        fmt.Println(v5, v6, v5)
                        _ = v41
                    } else {
                        switch (v5 + v6) {
                        case "":
                            var v42 int32 = -319
                            _ = v42
                        case "hello":
                            delete(v2, 127)
                        case "a":
                            if ((1 == 42) || false) {
                                fmt.Println(v5, v5)
                            }
                        }
                        v43 := GenS2{F0: false}
                        v5 = (v5 + (("\n\t" + "") + v5))
                        _ = v43
                    }
                    v1[""] = v6
                    var v44 bool = (true || true)
                    _ = v44
                }
            }
            var v45 int8 = 127
            v45 = genFn0()
            _ = v45
        }
        _ = fi34
    }
    v6 = (v5 + (((v5 + v5) + ("\n\t" + v5)) + "hello"))
    if ((127 != -297) && true) {
        fmt.Println(v6, v5)
        v5 = v6
        v46 := GenS3{F0: (genFn0() < -1), F1: 518, F2: 65535}
        _ = v46
    }
    var v47 bool = true
    v48 := []int{}
    v47 = ((4294967295 < 97) || (!(!true)))
    var v49 uint = 477
    v6 = ("hello" + v6)
    v50 := GenS2{F0: (!v47)}
    v49 = v49
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
