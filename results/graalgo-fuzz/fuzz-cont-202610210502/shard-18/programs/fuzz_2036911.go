package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 int; F2 int16 }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8; F1 int64 }
type GenS7 struct { F0 uint32 }
func genFn0(p0 uint8, p1 uint8) int {
    v6 := []int{0}
    var v7 uint32 = 2147483648
    fmt.Println(v7)
    v7 -= (v7 + v7)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 string, p1 int64) int8 {
    p0 = ((p0 + ("hello" + p0)) + p0)
    v8 := GenS1{F0: 529}
    var v9 int = 4294967295
    for fi10 := 0; fi10 < 8; fi10++ {
        if (false || (!(1 != 307))) {
            fmt.Println(p0, fi10)
            for fi11 := 0; fi11 < 19; fi11++ {
                fmt.Println(p1, fi10)
                var v12 int = fi11
                switch "" {
                case "foo bar":
                    v13 := GenS1{F0: 415}
                    _ = v13
                case "hello":
                    v14 := []int{}
                    _ = v14
                }
                _ = fi11
                _ = v12
            }
        } else {
            switch -826 {
            case 42:
                if false {
                    var v15 int = -448
                    var v16 uint64 = 999
                    _ = v15
                    _ = v16
                } else {
                    var v17 int64 = p1
                    if (!true) {
                        var v18 int32 = 255
                        if false {
                            var v19 int = 42
                            for fi20 := 0; fi20 < 7; fi20++ {
                                if (3.14159 <= (1.5 / 1.5)) {
                                    switch fi10 {
                                    case 4294967295:
                                        v21 := GenS1{F0: 2147483648}
                                        _ = v21
                                    case -458:
                                        switch (fi20 + (v9 * genFn0(2, 0))) {
                                        case 943:
                                            var v22 int = (fi20 - fi20)
                                            _ = v22
                                        case 1:
                                            var v23 int16 = 32767
                                            _ = v23
                                        }
                                    case -1:
                                        v24 := []int{4294967295, fi20, genFn0(156, 210), (-(fi10))}
                                        _ = v24
                                    }
                                    v17 += v17
                                    var v25 int16 = -64
                                    var v26 bool = true
                                    _ = v25
                                    _ = v26
                                } else {
                                    v9 -= v19
                                    v27 := GenS1{F0: 4294967295}
                                    fmt.Println(v19, v9)
                                    _ = v27
                                }
                                v28 := GenS2{F0: v17}
                                v19 += genFn0(127, 128)
                                _ = fi20
                                _ = v28
                            }
                            _ = v19
                        } else {
                            v9 -= fi10
                        }
                        v29 := GenS2{F0: 42}
                        _ = v18
                        _ = v29
                    }
                    _ = v17
                }
            case -300:
                v30 := GenS3{F0: 2, F1: (fi10 | 32768), F2: -1}
                _ = v30
            default:
                for fi31 := 0; fi31 < 12; fi31++ {
                    var v32 string = "foo bar"
                    _ = fi31
                    _ = v32
                }
            }
            var v33 float32 = 1e10
            var v34 uint = 2
            _ = v33
            _ = v34
        }
        if (296 > 908) {
            v35 := GenS1{F0: 4294967295}
            var v36 string = (p0 + "foo bar")
            v37 := GenS2{F0: (p1 >> ((1) & 63))}
            _ = v35
            _ = v36
            _ = v37
        }
        v38 := GenS3{F0: 1, F1: genFn0(140, 0), F2: 1}
        _ = fi10
        _ = v38
    }
    v39 := GenS4{F0: 18446744073709551615, F1: 65535}
    _ = p0
    _ = p1
    _ = v39
    _ = v8
    _ = v9
    return -92
}

func main() {
    var v1 bool = (18 > 16)
    v2 := make(map[string]int)
    var v3 uint8 = 1
    v4 := []uint32{1, 37, 533}
    v5 := make(map[string]string)
    if ((701 != -1) || false) {
        if ("a" == ("foo bar" + "")) {
            v3 -= v3
            v3 += (v3 + (v3 * v3))
            for fi40 := 0; fi40 < 10; fi40++ {
                v41 := GenS5{F0: -396}
                delete(v5, ("hello" + ""))
                fmt.Println(fi40, fi40)
                _ = fi40
                _ = v41
            }
        } else {
            v3 += v3
            if v1 {
                var v42 string = "foo bar"
                _ = v42
            }
            var v43 uint64 = 0
            _ = v43
        }
        for fi44 := 0; fi44 < 12; fi44++ {
            for fi45 := 0; fi45 < 10; fi45++ {
                v1 = false
                v46 := GenS3{F0: 1, F1: (fi44 + fi44), F2: 0}
                _ = fi45
                _ = v46
            }
            _ = fi44
        }
        v47 := v2[""]
        _ = v47
    }
    v3 -= (v3 / (((v3 - (v3 << ((46) & 63)))) | 1))
    v48 := GenS2{F0: 155}
    _ = len(v4)
    var v49 bool = (144 == 42)
    switch "a" {
    case "":
        for fi50 := 0; fi50 < 9; fi50++ {
            v51 := GenS6{F0: 100, F1: 1}
            if true {
                v49 = (32767 > 294)
                var v52 int = (-(fi50))
                _ = v52
            }
            _ = fi50
            _ = v51
        }
    }
    v49 = (2147483647 > -119)
    var v53 bool = (2147483648 >= 65535)
    if (1e10 <= 1e10) {
        v54 := GenS7{F0: 1}
        _ = v54
    } else {
        for fi55 := 0; fi55 < 10; fi55++ {
            if v53 {
                switch (("\n\t" + ("hello" + "hello")) + "foo bar") {
                case "\n\t":
                    v53 = true
                }
                var v56 string = ""
                var v57 float32 = 1e10
                _ = v56
                _ = v57
            }
            _ = fi55
        }
        var v58 uint8 = (^(v3))
        _ = v58
    }
    switch genFn0(50, v3) {
    case 4294967295:
        var v59 int = 1
        _ = v59
    }
    v60 := v2["\n\t"]
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v53
    _ = v60
}
