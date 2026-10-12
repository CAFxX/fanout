package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint; F1 int16; F2 uint32 }
type GenS6 struct { F0 int; F1 int; F2 uint32 }
func genFn0(p0 int32, p1 string) bool {
    for fi7 := 0; fi7 < 20; fi7++ {
        p0 -= 154
        var v8 string = ((p1 + "hello") + p1)
        for fi9 := 0; fi9 < 16; fi9++ {
            if (255 != 2) {
                var v10 uint64 = 4294967295
                _ = v10
            }
            v8 = ""
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        switch fi11 {
        case 272:
            var v12 uint8 = 128
            _ = v12
        case 756:
            for fi13 := 0; fi13 < 5; fi13++ {
                v14 := GenS1{F0: 1, F1: -32768, F2: 380}
                if (!true) {
                    var v15 int32 = (p0 - p0)
                    v16 := GenS1{F0: 1, F1: 1000, F2: 823}
                    _ = v15
                    _ = v16
                }
                var v17 uint = 599
                var v18 int32 = 1
                if ((!true) || (0 >= 568)) {
                    v18 = 42
                    switch (p1 + (("a" + (((p1 + (p1 + (("foo bar" + "") + (("\n\t" + p1) + p1)))) + p1) + (p1 + p1))) + p1)) {
                    case "a":
                        var v19 float32 = 3.14159
                        _ = v19
                    case "":
                        var v20 uint8 = 128
                        _ = v20
                    }
                } else {
                    var v21 int32 = 2147483647
                    if (2147483648 < 2147483648) {
                        fmt.Println(v21, v21)
                        switch p1 {
                        case "\n\t":
                            fmt.Println(v17, p1, p1)
                        case "foo bar":
                            v22 := []int{}
                            _ = v22
                        }
                        v23 := GenS2{F0: 794}
                        var v24 bool = true
                        _ = v23
                        _ = v24
                    }
                    v21 += (v21 ^ (p0 / ((v18) | 1)))
                    _ = v21
                }
                _ = fi13
                _ = v14
                _ = v17
                _ = v18
            }
        }
        _ = fi11
    }
    p1 = (("a" + (p1 + "")) + p1)
    v25 := []int{-784, -966, 48, -2147483648}
    var v26 int16 = 0
    _ = p0
    _ = p1
    _ = v25
    _ = v26
    return true
}
func genFn1() uint64 {
    if (true || (598 != -600)) {
        var v27 bool = (117 != 42)
        for fi28 := 0; fi28 < 8; fi28++ {
            v27 = (42 != -1)
            if (!(1 != 4294967295)) {
                v27 = (42 > -128)
            }
            _ = fi28
        }
        _ = v27
    } else {
        for fi29 := 0; fi29 < 8; fi29++ {
            v30 := []int{}
            var v31 uint = 229
            _ = fi29
            _ = v30
            _ = v31
        }
        v32 := GenS2{F0: -32768}
        _ = v32
    }
    if true {
        v33 := GenS2{F0: 682}
        _ = v33
    }
    return 4294967295
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]bool)
    var v4 int8 = -1
    var v5 int64 = v1
    v6 := make(map[uint32]int)
    v34 := GenS3{F0: (-171 >= 1)}
    v5 = 0
    fmt.Println(v4)
    v35 := GenS4{F0: 2147483647}
    v4 += v4
    delete(v3, 830)
    var v36 int = 659
    if (true && (!(v5 < 9223372036854775807))) {
        if (463 > -297) {
            v37 := GenS2{F0: 409}
            v38 := GenS2{F0: 42}
            _ = v37
            _ = v38
        }
    } else {
        v36 -= v36
    }
    if true {
        for fi39 := 0; fi39 < 8; fi39++ {
            v5 = -260
            _ = fi39
        }
    }
    for fi40 := 0; fi40 < 20; fi40++ {
        for fi41 := 0; fi41 < 15; fi41++ {
            var v42 int64 = (v1 & -804)
            _ = fi41
            _ = v42
        }
        if true {
            v43 := GenS3{F0: true}
            v4 += -128
            for fi44 := 0; fi44 < 19; fi44++ {
                var v45 uint64 = 1
                v45 = (v45 * 0)
                _ = fi44
                _ = v45
            }
            _ = v43
        } else {
            v46 := GenS4{F0: 461}
            _ = v46
        }
        _ = fi40
    }
    var v47 int32 = 0
    v48 := GenS5{F0: 0, F1: 1000, F2: 4294967295}
    _ = len(v2)
    var v49 uint8 = 1
    for fi50 := 0; fi50 < 16; fi50++ {
        v51 := GenS6{F0: fi50, F1: 128, F2: 0}
        _ = fi50
        _ = v51
    }
    var v52 int64 = (v1 % (((v5 + -188)) | 1))
    fmt.Println(v49, v36)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v52
    _ = v6
}
