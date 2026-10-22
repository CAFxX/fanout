package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 string; F1 int; F2 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint8, p1 uint64) bool {
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := GenS1{F0: 191, F1: 100}
        p1 = 1
        _ = fi6
        _ = v7
    }
    var v8 int32 = -1
    v9 := []int{1, -2147483648, -747, -15}
    fmt.Println(p0, p1, v8)
    switch 32768 {
    case 9223372036854775807:
        if (614 < uint32(p1)) {
            var v10 uint32 = 2147483648
            fmt.Println(p1, p1)
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = v10
        } else {
            switch -9223372036854775808 {
            case -2147483648:
                for fi11 := 0; fi11 < 3; fi11++ {
                    p0 = p0
                    var v12 int = fi11
                    _ = fi11
                    _ = v12
                }
            case 127:
                for fi13 := 0; fi13 < 18; fi13++ {
                    if true {
                        fmt.Println(fi13, p1)
                    } else {
                        var v14 uint16 = 361
                        _ = v14
                    }
                    _ = fi13
                }
            }
            if true {
                switch 1 {
                case 9223372036854775807:
                    p1 = p1
                }
                var v15 int = -12
                _ = v15
            }
            v16 := GenS1{F0: uint(p0), F1: 127}
            _ = v16
        }
    case 809:
        var v17 uint8 = 0
        _ = v17
    default:
        if false {
            var v18 int32 = v8
            _ = v18
        }
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return false
}
func genFn1() int16 {
    v19 := GenS2{F0: "", F1: -9223372036854775808, F2: 255}
    for fi20 := 0; fi20 < 1; fi20++ {
        v21 := GenS2{F0: ("hello" + "hello"), F1: (fi20 ^ (^(fi20))), F2: 2}
        v22 := []int{(fi20 / ((4294967295) | 1)), fi20, (fi20 * (fi20 & 0))}
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 uint16 = 842
    fmt.Println(v23)
    _ = v19
    _ = v23
    return 0
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[uint32]bool)
    var v3 bool = true
    var v4 int32 = -1
    var v5 float32 = ((3.14159 + (-2.25 + 1e10)) + 1.5)
    v24 := GenS3{F0: 0}
    _ = len(v2)
    var v25 uint8 = (v1 % (((v1 % (((v1 % (((v1 | v1)) | 1))) | 1))) | 1))
    fmt.Println(v5, v5)
    if true {
        v26 := GenS4{F0: v4}
        _ = v26
    } else {
        if v3 {
            v27 := []int{2147483647, -2147483648}
            _ = v27
        }
        var v28 uint16 = 42
        var v29 uint = 550
        _ = v28
        _ = v29
    }
    var v30 uint = 726
    fmt.Println(v25)
    for fi31 := 0; fi31 < 3; fi31++ {
        for fi32 := 0; fi32 < 1; fi32++ {
            delete(v2, 598)
            var v33 int = (fi31 - (fi32 - 32767))
            var v34 uint32 = 42
            _ = fi32
            _ = v33
            _ = v34
        }
        fmt.Println(v25, v1)
        _ = fi31
    }
    if v3 {
        for fi35 := 0; fi35 < 2; fi35++ {
            for fi36 := 0; fi36 < 5; fi36++ {
                var v37 bool = (!v3)
                v25 = (v1 / ((v1) | 1))
                _ = fi36
                _ = v37
            }
            _ = fi35
        }
        v38 := GenS4{F0: v4}
        _ = v38
    } else {
        var v39 bool = v3
        var v40 uint = (v30 % ((v30) | 1))
        v39 = (198 <= 42)
        _ = v39
        _ = v40
    }
    for fi41 := 0; fi41 < 18; fi41++ {
        v42 := GenS1{F0: (^(v30)), F1: -128}
        var v43 uint32 = 4294967295
        _ = fi41
        _ = v42
        _ = v43
    }
    var v44 uint32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v44
    _ = v5
}
