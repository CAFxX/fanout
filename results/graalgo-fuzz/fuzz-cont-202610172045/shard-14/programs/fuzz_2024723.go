package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 uint32 }
func genFn0() bool {
    if false {
        if false {
            var v6 uint32 = 0
            _ = v6
        } else {
            v7 := []int{9223372036854775807, -761, 127}
            var v8 uint8 = 21
            v9 := GenS1{F0: 1}
            _ = v7
            _ = v8
            _ = v9
        }
    } else {
        var v10 string = "a"
        v11 := []int{9223372036854775807, 42, 2147483647}
        v12 := GenS1{F0: -116}
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 int = -9223372036854775808
    v13 -= v13
    _ = v13
    return false
}
func genFn1(p0 int8, p1 int32) string {
    var v14 int = -241
    for fi15 := 0; fi15 < 11; fi15++ {
        if true {
            var v16 int32 = p1
            var v17 bool = false
            v16 = 0
            _ = v16
            _ = v17
        }
        for fi18 := 0; fi18 < 3; fi18++ {
            v19 := GenS1{F0: 649}
            v20 := GenS1{F0: 0}
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi15
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        if false {
            v14 = (v14 * (fi21 >> ((2) & 63)))
            switch (fi21 % (((v14 * fi21)) | 1)) {
            case 128:
                v22 := []int{v14, -485, (fi21 - 2147483647), v14}
                _ = v22
            case -1:
                v23 := []int{(^(fi21)), v14}
                _ = v23
            }
            var v24 uint8 = 127
            _ = v24
        } else {
            v25 := GenS2{F0: 4294967295, F1: 1}
            v26 := GenS1{F0: -1}
            if true {
                var v27 uint32 = 1
                _ = v27
            } else {
                v14 = (fi21 - 2147483647)
                v28 := []int{9223372036854775807, fi21}
                _ = v28
            }
            _ = v25
            _ = v26
        }
        for fi29 := 0; fi29 < 2; fi29++ {
            var v30 uint64 = 1
            v31 := []int{fi21, fi29, 32767, 32767}
            _ = fi29
            _ = v30
            _ = v31
        }
        v32 := []int{(v14 - (v14 - (v14 | (fi21 * (fi21 - fi21)))))}
        _ = fi21
        _ = v32
    }
    v33 := GenS3{F0: 1}
    p1 = p1
    _ = p0
    _ = p1
    _ = v14
    _ = v33
    return ""
}

func main() {
    var v1 string = "foo bar"
    var v2 bool = ((-128 >= -128) || false)
    v3 := make(map[int]int)
    v4 := make(map[string]bool)
    var v5 float32 = float32(3.14159)
    v34 := GenS4{F0: 42, F1: -59}
    var v35 bool = true
    if v2 {
        v36 := GenS3{F0: 65535}
        fmt.Println(v2)
        var v37 string = "foo bar"
        _ = v36
        _ = v37
    } else {
        v35 = false
        v5 = (3.14159 - v5)
    }
    if (!(!v35)) {
        if (1000 >= 794) {
            v35 = v35
            var v38 int32 = 535
            var v39 int64 = -1
            _ = v38
            _ = v39
        }
        if v2 {
            v1 = ""
            v40 := GenS5{F0: 2}
            _ = v40
        }
    }
    var v41 uint = 42
    if (!v35) {
        var v42 int16 = 32767
        v3[-763] = 128
        _ = v42
    }
    if v2 {
        v43 := []int{32767, -976, 128}
        _ = v43
    }
    var v44 uint16 = 65535
    var v45 float32 = -2.25
    v5 -= ((0.1 + 3.14159) - (v45 - v5))
    switch (((v1 + ("a" + "hello")) + v1) + (v1 + (v1 + v1))) {
    case "\n\t":
        var v46 int32 = 994
        _ = v46
    case "foo bar":
        for fi47 := 0; fi47 < 17; fi47++ {
            if true {
                var v48 uint = (v41 / ((v41) | 1))
                _ = v48
            } else {
                v45 = v45
                v2 = (42 != -56)
                v2 = (127 == 127)
            }
            v2 = v2
            v1 = (v1 + v1)
            _ = fi47
        }
    case "hello":
        var v49 int8 = 100
        _ = v49
    default:
        v1 = v1
    }
    if v2 {
        for fi50 := 0; fi50 < 12; fi50++ {
            var v51 int32 = 1
            var v52 int16 = -1
            _ = fi50
            _ = v51
            _ = v52
        }
        v53 := GenS1{F0: -1}
        _ = v53
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v41
    _ = v44
    _ = v45
    _ = v5
}
