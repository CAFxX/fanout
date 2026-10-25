package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 string; F2 bool }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 int32; F1 string; F2 string }
func genFn0() int32 {
    var v6 uint = 18446744073709551615
    for fi7 := 0; fi7 < 3; fi7++ {
        if (true || false) {
            var v8 uint32 = 42
            v8 -= v8
            _ = v8
        } else {
            for fi9 := 0; fi9 < 13; fi9++ {
                var v10 uint16 = 42
                v11 := GenS1{F0: 127}
                if true {
                    var v12 uint16 = (v10 & v10)
                    var v13 int = fi9
                    _ = v12
                    _ = v13
                }
                _ = fi9
                _ = v10
                _ = v11
            }
        }
        v14 := GenS1{F0: 0}
        var v15 uint8 = 255
        _ = fi7
        _ = v14
        _ = v15
    }
    _ = v6
    return 1
}
func genFn1(p0 int32) uint64 {
    var v16 uint32 = 478
    var v17 uint64 = 4294967295
    fmt.Println(v17)
    _ = p0
    _ = v16
    _ = v17
    return v17
}

func main() {
    var v1 int32 = 2147483647
    var v2 int = -2147483648
    var v3 int = v2
    v4 := make(map[string]string)
    var v5 bool = true
    if v5 {
        _ = len(v4)
        fmt.Println(v3, v3)
    }
    var v18 int64 = -1
    switch (v3 - -40) {
    case 4294967295:
        fmt.Println(v3)
    }
    fmt.Println(v5, v3)
    for fi19 := 0; fi19 < 14; fi19++ {
        for fi20 := 0; fi20 < 10; fi20++ {
            if (!(!v5)) {
                v21 := GenS2{F0: 805}
                if v5 {
                    v22 := GenS3{F0: 0, F1: "hello", F2: v5}
                    var v23 int8 = 127
                    switch (fi20 >> ((42) & 63)) {
                    case -1:
                        if ((32767 > 905) && (!(v1 <= v1))) {
                            v24 := GenS2{F0: 1000}
                            _ = v24
                        } else {
                            v18 = (v18 % (((v18 | (v18 % ((v18) | 1)))) | 1))
                            _ = len(v4)
                        }
                    default:
                        var v25 int = (fi20 / (((v3 | fi20)) | 1))
                        _ = v25
                    }
                    fmt.Println(v3, fi20, v3)
                    _ = v22
                    _ = v23
                } else {
                    fmt.Println(v5, v5, v18)
                    v26 := []int{v3, -9223372036854775808}
                    v1 -= v1
                    _ = v26
                }
                _ = v21
            } else {
                var v27 string = "\n\t"
                _ = v27
            }
            v28 := []int{(-(v2)), (^(v2))}
            v4[""] = "\n\t"
            _ = fi20
            _ = v28
        }
        var v29 string = "a"
        _ = fi19
        _ = v29
    }
    for fi30 := 0; fi30 < 11; fi30++ {
        v3 -= (fi30 >> ((4294967295) & 63))
        _ = fi30
    }
    var v31 uint = 807
    var v32 uint16 = 415
    if false {
        var v33 float32 = 3.14159
        var v34 float32 = 0.0
        v35 := GenS4{F0: v5, F1: (v33 != v33)}
        _ = v33
        _ = v34
        _ = v35
    }
    v36 := GenS5{F0: v1, F1: ("foo bar" + ("\n\t" + ("a" + "foo bar"))), F2: ("hello" + "hello")}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v36
    _ = v4
    _ = v5
}
