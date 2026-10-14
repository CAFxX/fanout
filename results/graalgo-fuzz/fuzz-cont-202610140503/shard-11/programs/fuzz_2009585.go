package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 int; F2 bool }
func genFn0(p0 int16, p1 bool) uint32 {
    var v5 uint = 179
    if p1 {
        v6 := []int{32768, -409, -9223372036854775808, -106}
        p0 = 0
        _ = v6
    }
    p1 = p1
    v7 := GenS1{F0: 18446744073709551615}
    _ = p0
    _ = p1
    _ = v5
    _ = v7
    return 259
}
func genFn1(p0 string, p1 uint64) int8 {
    p0 = (p0 + (((p0 + p0) + (p0 + "foo bar")) + "hello"))
    fmt.Println(p0, p1)
    var v8 int32 = 163
    p1 = (p1 % ((1) | 1))
    _ = p0
    _ = p1
    _ = v8
    return 31
}

func main() {
    var v1 string = "hello"
    var v2 bool = true
    v3 := make(map[string]bool)
    var v4 string = v1
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 uint32 = 2
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 4294967295}
    switch 1 {
    case 127:
        var v12 uint8 = 127
        _ = v12
    case 42:
        fmt.Println(v2)
    case 0:
        for fi13 := 0; fi13 < 5; fi13++ {
            v1 = v4
            if true {
                var v14 bool = v2
                var v15 string = v1
                var v16 float32 = 0.1
                _ = v14
                _ = v15
                _ = v16
            } else {
                for fi17 := 0; fi17 < 4; fi17++ {
                    v1 = v4
                    _ = fi17
                }
            }
            v1 = v1
            _ = fi13
        }
    }
    if (-1 == 316) {
        var v18 bool = v2
        var v19 uint16 = 2
        if false {
            switch -346 {
            case 42:
                switch 32767 {
                case -479:
                    var v20 int = 0
                    _ = v20
                case 20:
                    fmt.Println(v18)
                }
            case -9223372036854775808:
                switch 743 {
                case -1:
                    _ = len(v3)
                case -2147483648:
                    for fi21 := 0; fi21 < 16; fi21++ {
                        fmt.Println(fi21, v4, v1)
                        _ = fi21
                    }
                case 4294967295:
                    v22 := []int{9223372036854775807, 32768, int(v19)}
                    _ = v22
                }
            case -1:
                delete(v3, v1)
            default:
                fmt.Println(v1)
            }
            for fi23 := 0; fi23 < 3; fi23++ {
                fmt.Println(v19, fi23)
                if false {
                    for fi24 := 0; fi24 < 18; fi24++ {
                        if (true && false) {
                            var v25 uint64 = 0
                            _ = v25
                        }
                        v26 := GenS2{F0: 127, F1: 4294967295, F2: 1}
                        _ = fi24
                        _ = v26
                    }
                }
                _ = fi23
            }
        } else {
            v3["hello"] = ((!true) || false)
            v1 = v1
        }
        _ = v18
        _ = v19
    }
    for fi27 := 0; fi27 < 12; fi27++ {
        v28 := GenS3{F0: 42}
        var v29 uint8 = 0
        v1 = (v1 + (v1 + "a"))
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 uint32 = 2
    if (v1 >= v4) {
        v31 := GenS1{F0: 668}
        var v32 uint16 = 0
        _ = v31
        _ = v32
    }
    if v2 {
        v33 := GenS1{F0: 2}
        _ = v33
    } else {
        v1 = "hello"
        for fi34 := 0; fi34 < 5; fi34++ {
            v2 = v2
            v35 := GenS4{F0: 414, F1: 32768, F2: ((v30 > v30) && false)}
            v36 := []int{fi34, 128, fi34}
            _ = fi34
            _ = v35
            _ = v36
        }
        for fi37 := 0; fi37 < 3; fi37++ {
            for fi38 := 0; fi38 < 2; fi38++ {
                var v39 int64 = 9223372036854775807
                _ = fi38
                _ = v39
            }
            _ = fi37
        }
    }
    v40 := GenS1{F0: 614}
    for fi41 := 0; fi41 < 1; fi41++ {
        v42 := []int{fi41, fi41, fi41}
        _ = fi41
        _ = v42
    }
    var v43 int = 380
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v40
    _ = v43
}
