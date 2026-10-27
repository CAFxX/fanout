package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 string; F1 int; F2 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 bool; F1 bool }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int32; F1 uint; F2 int16 }
type GenS9 struct { F0 int; F1 uint; F2 int32 }
type GenS10 struct { F0 int16; F1 int16; F2 bool }
func genFn0(p0 int16, p1 uint8) uint32 {
    var v7 int32 = -2147483648
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := []int{(^(fi8)), fi8, 127, (fi8 / ((fi8) | 1))}
        fmt.Println(v7)
        _ = fi8
        _ = v9
    }
    switch (("\n\t" + "a") + "hello") {
    case "\n\t":
        v10 := GenS1{F0: "\n\t", F1: 974}
        _ = v10
    }
    v11 := GenS2{F0: "", F1: 576, F2: 633}
    for fi12 := 0; fi12 < 14; fi12++ {
        v13 := []int{-1, int(p1)}
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    return 0
}
func genFn1(p0 uint32) int8 {
    var v14 uint32 = genFn0(918, 255)
    if false {
        for fi15 := 0; fi15 < 16; fi15++ {
            v16 := GenS3{F0: false}
            _ = fi15
            _ = v16
        }
        v17 := GenS4{F0: 29, F1: (v14 / (((v14 % ((42) | 1))) | 1))}
        var v18 uint16 = 65535
        _ = v17
        _ = v18
    } else {
        var v19 uint32 = p0
        v20 := []int{-1, 428, 42}
        fmt.Println(v14, p0)
        _ = v19
        _ = v20
    }
    var v21 int8 = -1
    if (4294967295 == 4294967295) {
        v22 := []int{-189, 0, -614, -9223372036854775808}
        _ = v22
    }
    _ = p0
    _ = v14
    _ = v21
    return v21
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = "foo bar"
    v3 := make(map[uint32]bool)
    var v4 float32 = 0.0
    var v5 uint64 = 873
    var v6 int16 = 339
    v23 := GenS2{F0: v2, F1: 0, F2: 758}
    if false {
        for fi24 := 0; fi24 < 1; fi24++ {
            for fi25 := 0; fi25 < 20; fi25++ {
                switch v2 {
                case "\n\t":
                    var v26 int16 = v6
                    _ = v26
                case "hello":
                    if (true && (128 > 1)) {
                        v2 = ("foo bar" + (v2 + "hello"))
                    } else {
                        switch fi25 {
                        case 32768:
                            v5 = v5
                        case 933:
                            v27 := GenS3{F0: true}
                            _ = v27
                        case 848:
                            var v28 bool = (((!true) || ((0.0 + 1.5) <= 0.1)) || (2 > 0))
                            var v29 uint64 = 839
                            _ = v28
                            _ = v29
                        }
                        var v30 uint8 = 127
                        var v31 int8 = -1
                        _ = v30
                        _ = v31
                    }
                    v5 += 2
                }
                var v32 int64 = 0
                if false {
                    v33 := GenS1{F0: (v2 + ""), F1: (fi24 ^ fi25)}
                    var v34 float32 = (((v4 + ((v4 * float32(1e10)) + v4)) + float32(float64(float32(1.5)))) / (v4 / v4))
                    var v35 uint8 = 2
                    _ = v33
                    _ = v34
                    _ = v35
                }
                if (true || (1.5 != (((((0.1 / (v4 + v4)) / 1.5) - (v4 - v4)) - (v4 + ((v4 + 3.14159) / -2.25))) - float32((float64(float32((((0.0 / ((1e10 / 1.5) + 0.1)) * 0.1) * 1e10))) * 1e10))))) {
                    if false {
                        var v36 float64 = 0.0
                        _ = v36
                    }
                }
                _ = fi25
                _ = v32
            }
            v37 := GenS2{F0: v2, F1: 127, F2: 2147483648}
            switch fi24 {
            case 128:
                if (("\n\t" + "hello") != "a") {
                    v38 := GenS5{F0: (fi24 - 2147483647)}
                    var v39 int8 = -1
                    v40 := GenS6{F0: ((v2 + ((("foo bar" + v2) + "hello") + ("hello" + v2))) <= v2), F1: true}
                    _ = v38
                    _ = v39
                    _ = v40
                }
            case 0:
                var v41 float32 = v4
                _ = v41
            case 698:
                v42 := GenS7{F0: 1}
                _ = v42
            }
            _ = fi24
            _ = v37
        }
        if false {
            v43 := GenS6{F0: true, F1: false}
            v2 = "a"
            for fi44 := 0; fi44 < 16; fi44++ {
                v45 := GenS8{F0: 0, F1: 255, F2: (v6 >> ((1) & 63))}
                fmt.Println(v4, v6, v4)
                var v46 string = (("foo bar" + (v2 + (v2 + (v2 + (v2 + v2))))) + "a")
                _ = fi44
                _ = v45
                _ = v46
            }
            _ = v43
        } else {
            var v47 int32 = 2147483647
            _ = v47
        }
        var v48 string = v2
        _ = v48
    }
    if (!(!(-798 > -1))) {
        v2 = v2
        var v49 uint32 = 2147483648
        _ = v49
    }
    var v50 uint16 = 0
    var v51 bool = true
    if v51 {
        delete(v3, 4294967295)
    } else {
        v52 := GenS9{F0: -2147483648, F1: 18446744073709551615, F2: -2147483648}
        fmt.Println(v2)
        v2 = ""
        _ = v52
    }
    v53 := []int{9223372036854775807, -9223372036854775808, 2147483647}
    fmt.Println(v6, v5)
    v50 = v50
    var v54 string = "hello"
    v6 = v6
    if (0 < 1) {
        var v55 uint8 = 46
        for fi56 := 0; fi56 < 14; fi56++ {
            fmt.Println(fi56, fi56, v55)
            var v57 uint32 = 471
            _ = fi56
            _ = v57
        }
        _ = v55
    }
    v58 := GenS10{F0: 853, F1: (v6 >> ((18446744073709551615) & 63)), F2: (1 != 758)}
    var v59 bool = (v6 >= 0)
    var v60 string = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v53:", len(v53), v53)
    fmt.Println("v54:", v54)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v53
    _ = v54
    _ = v58
    _ = v59
    _ = v6
    _ = v60
}
