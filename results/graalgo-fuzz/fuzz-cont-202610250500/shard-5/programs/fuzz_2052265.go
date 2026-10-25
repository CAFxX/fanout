package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }
func genFn0() int64 {
    var v6 string = ("foo bar" + "")
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := []int{}
        _ = fi7
        _ = v8
    }
    _ = v6
    return 0
}
func genFn1() bool {
    var v9 uint32 = 88
    v9 += 0
    for fi10 := 0; fi10 < 9; fi10++ {
        fmt.Println(v9, v9)
        if false {
            if (2 < 255) {
                var v11 string = ("foo bar" + "")
                var v12 uint64 = 42
                _ = v11
                _ = v12
            } else {
                var v13 uint64 = 4294967295
                var v14 uint = 370
                _ = v13
                _ = v14
            }
            if false {
                v15 := []int{-692, fi10}
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                for fi16 := 0; fi16 < 17; fi16++ {
                    var v17 uint32 = v9
                    var v18 uint = 318
                    v18 += v18
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                _ = v15
            } else {
                for fi19 := 0; fi19 < 2; fi19++ {
                    var v20 string = "\n\t"
                    _ = fi19
                    _ = v20
                }
            }
        }
        v9 += v9
        _ = fi10
    }
    _ = v9
    return (uint8(v9) != 42)
}

func main() {
    var v1 int16 = 333
    v2 := []int{}
    var v3 uint16 = 353
    v4 := make(map[uint32]string)
    v5 := []int{}
    if false {
        switch 9223372036854775807 {
        case 32767:
            v21 := GenS1{F0: 539}
            _ = v21
        default:
            v1 += -116
        }
    }
    if (v1 < 0) {
        for fi22 := 0; fi22 < 3; fi22++ {
            switch fi22 {
            case 655:
                fmt.Println(v1, v3)
            }
            for fi23 := 0; fi23 < 3; fi23++ {
                v24 := GenS1{F0: uint(v3)}
                v25 := GenS1{F0: 255}
                v26 := GenS1{F0: 255}
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v27 := GenS1{F0: 4294967295}
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
                _ = v27
            }
            v3 += (v3 | v3)
            _ = fi22
        }
        _ = len(v5)
    } else {
        var v28 uint16 = 417
        v28 = (v3 + (v28 & (v3 * v3)))
        v1 += v1
        _ = v28
    }
    var v29 int8 = 64
    var v30 uint = 18446744073709551615
    switch 475 {
    case 0:
        switch 4294967295 {
        case -968:
            v31 := GenS2{F0: (v30 ^ (v30 << (((v30 | 42)) & 63)))}
            _ = v31
        default:
            if (true || false) {
                if false {
                    fmt.Println(v30, v1, v1)
                    var v32 uint64 = 829
                    _ = v32
                } else {
                    _ = len(v5)
                }
            } else {
                var v33 string = (((("\n\t" + ((("" + "a") + ("a" + ((("\n\t" + ("\n\t" + (("\n\t" + "") + "\n\t"))) + (("foo bar" + ("foo bar" + "")) + "hello")) + "foo bar"))) + "")) + "\n\t") + "\n\t") + "foo bar")
                _ = v33
            }
        }
    case 128:
        for fi34 := 0; fi34 < 6; fi34++ {
            if false {
                v30 = (^(v30))
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi34
        }
    }
    for fi35 := 0; fi35 < 20; fi35++ {
        var v36 int32 = -1
        _ = fi35
        _ = v36
    }
    v37 := GenS3{F0: -2147483648}
    var v38 uint64 = 793
    delete(v4, 171)
    v1 += v1
    v38 = 42
    for fi39 := 0; fi39 < 11; fi39++ {
        if ((-2.25 <= 0.0) || true) {
            var v40 uint = v30
            switch -470 {
            case 128:
                v30 = 935
            case 2147483647:
                var v41 uint32 = uint32(v3)
                _ = v41
            }
            _ = v40
        }
        _ = fi39
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
