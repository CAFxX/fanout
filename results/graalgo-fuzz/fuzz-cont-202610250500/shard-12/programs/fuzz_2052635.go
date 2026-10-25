package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 uint8; F2 string }
type GenS4 struct { F0 bool }
func genFn0() int32 {
    var v7 int16 = 1000
    if false {
        v7 += (v7 & v7)
        if (658 >= 65535) {
            v8 := []int{}
            if ((42 < 127) && true) {
                var v9 uint16 = 42
                for fi10 := 0; fi10 < 13; fi10++ {
                    v11 := GenS1{F0: ((("foo bar" + "hello") + "\n\t") + "hello"), F1: 255}
                    v8 = append(v8, (fi10 << ((65535) & 63)))
                    _ = fi10
                    _ = v11
                }
                v9 = v9
                _ = v9
            }
            _ = v8
        } else {
            var v12 int16 = v7
            var v13 int16 = v7
            _ = v12
            _ = v13
        }
        v7 = (v7 * v7)
    } else {
        var v14 uint64 = 2
        v15 := []int{2147483647, 42, 937, 624}
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        fmt.Println(fi16)
        v17 := GenS2{F0: -749}
        _ = fi16
        _ = v17
    }
    var v18 bool = true
    v18 = true
    _ = v18
    _ = v7
    return -1
}
func genFn1() int64 {
    var v19 float32 = (1.5 - 0.1)
    for fi20 := 0; fi20 < 1; fi20++ {
        for fi21 := 0; fi21 < 8; fi21++ {
            v22 := GenS1{F0: ("\n\t" + "foo bar"), F1: 0}
            v23 := []int{42, (-(fi20)), (fi21 * (fi20 % (((fi20 * fi21)) | 1)))}
            v23 = append(v23, (fi20 % ((128) | 1)))
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi20
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        fmt.Println(fi24)
        fmt.Println(v19, fi24)
        for fi25 := 0; fi25 < 20; fi25++ {
            v19 -= 0.0
            var v26 uint64 = 2
            var v27 int64 = 9223372036854775807
            v28 := GenS3{F0: 782, F1: 156, F2: ""}
            switch (("foo bar" + "\n\t") + "") {
            case "foo bar":
                v26 = 0
            case "hello":
                v19 = -2.25
            default:
                v26 = v26
                v26 = v26
                v29 := GenS1{F0: "foo bar", F1: 0}
                _ = v29
            }
            _ = fi25
            _ = v26
            _ = v27
            _ = v28
        }
        _ = fi24
    }
    var v30 int64 = -9223372036854775808
    _ = v19
    _ = v30
    return v30
}

func main() {
    var v1 int16 = 198
    v2 := []uint32{2147483647}
    var v3 string = "foo bar"
    v4 := make(map[int]string)
    v5 := []string{""}
    var v6 uint16 = 65535
    v3 = (("foo bar" + v3) + v3)
    var v31 uint = 1
    v6 -= v6
    fmt.Println(v3, v3, v3)
    for fi32 := 0; fi32 < 10; fi32++ {
        v1 += 472
        if false {
            var v33 uint = v31
            var v34 float64 = 0.0
            _ = v33
            _ = v34
        } else {
            var v35 float64 = ((3.14159 + 3.14159) / 3.14159)
            v31 = 4294967295
            _ = v35
        }
        _ = fi32
    }
    delete(v4, 32768)
    delete(v4, -1)
    var v36 string = v3
    var v37 uint = v31
    for fi38 := 0; fi38 < 6; fi38++ {
        if (21 > -128) {
            v31 = (v31 % ((v31) | 1))
            switch 32768 {
            case 695:
                var v39 uint8 = 1
                _ = v39
            case 128:
                var v40 int = fi38
                _ = v40
            }
        } else {
            if true {
                _ = len(v4)
                v31 = (v37 ^ (^(v37)))
                var v41 uint = 1
                _ = v41
            } else {
                delete(v4, int(v37))
            }
        }
        v42 := GenS3{F0: v31, F1: 127, F2: "a"}
        _ = fi38
        _ = v42
    }
    var v43 uint8 = 195
    for fi44 := 0; fi44 < 5; fi44++ {
        fmt.Println(v1)
        v45 := GenS2{F0: 164}
        v4[1] = ""
        _ = fi44
        _ = v45
    }
    v6 = (v6 / ((v6) | 1))
    v46 := v4[32768]
    if ((42 == 32768) && false) {
        v47 := GenS4{F0: (1 >= 0)}
        var v48 bool = true
        var v49 int16 = v1
        _ = v47
        _ = v48
        _ = v49
    }
    var v50 float32 = 0.0
    _ = len(v5)
    var v51 float32 = (v50 + (v50 - ((1e10 + (0.0 * 0.0)) + (v50 + v50))))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v36
    _ = v37
    _ = v4
    _ = v43
    _ = v46
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
