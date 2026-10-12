package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int32 }
type GenS2 struct { F0 int; F1 int16; F2 bool }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 uint8 }
func genFn0() uint32 {
    var v7 float64 = 0.0
    fmt.Println(v7)
    _ = v7
    return 2147483648
}
func genFn1() int64 {
    v8 := GenS1{F0: "\n\t", F1: 1, F2: 1}
    for fi9 := 0; fi9 < 9; fi9++ {
        v10 := []int{-755, fi9}
        _ = fi9
        _ = v10
    }
    var v11 string = "foo bar"
    var v12 uint16 = 42
    for fi13 := 0; fi13 < 16; fi13++ {
        v12 += 2
        v14 := GenS1{F0: v11, F1: -76, F2: 42}
        _ = fi13
        _ = v14
    }
    _ = v11
    _ = v12
    _ = v8
    return -1
}

func main() {
    var v1 uint64 = 4294967295
    var v2 float32 = (3.14159 / 0.1)
    var v3 int64 = 934
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]uint32)
    v6 := []string{"", ("foo bar" + ((("foo bar" + ("" + "")) + "a") + "hello")), "hello", "foo bar"}
    if true {
        var v15 int32 = -2147483648
        v16 := GenS2{F0: -2147483648, F1: 1, F2: (uint32(v1) != 65535)}
        _ = v15
        _ = v16
    } else {
        fmt.Println(v3)
        if (false || (true || true)) {
            for fi17 := 0; fi17 < 1; fi17++ {
                v3 += (v3 * 244)
                v1 = (v1 % (((v1 - v1)) | 1))
                v3 -= (-(v3))
                _ = fi17
            }
        }
    }
    v18 := GenS3{F0: (127 >= -1), F1: 428}
    v1 += (v1 - v1)
    v19 := GenS4{F0: (("\n\t" + "") + "a"), F1: 2147483647}
    v5[242] = 0
    var v20 uint32 = 555
    if true {
        for fi21 := 0; fi21 < 7; fi21++ {
            var v22 uint32 = v20
            var v23 int64 = v3
            _ = fi21
            _ = v22
            _ = v23
        }
        v24 := GenS1{F0: ("" + "foo bar"), F1: 60, F2: -1}
        _ = v24
    }
    var v25 uint64 = (v1 - v1)
    v20 += genFn0()
    v25 = 18446744073709551615
    v26 := GenS4{F0: "a", F1: genFn0()}
    if true {
        fmt.Println(v3, v20, v20)
        for fi27 := 0; fi27 < 18; fi27++ {
            for fi28 := 0; fi28 < 1; fi28++ {
                fmt.Println(v20, v20, fi27)
                var v29 string = "a"
                _ = fi28
                _ = v29
            }
            if (834 == 0) {
                fmt.Println(v1)
            } else {
                v20 -= genFn0()
                for fi30 := 0; fi30 < 8; fi30++ {
                    v31 := GenS5{F0: 2}
                    var v32 uint64 = (v25 >> ((65535) & 63))
                    v33 := GenS4{F0: ((("\n\t" + "foo bar") + ("\n\t" + "\n\t")) + ("" + "foo bar")), F1: v20}
                    _ = fi30
                    _ = v31
                    _ = v32
                    _ = v33
                }
            }
            _ = fi27
        }
    }
    v2 -= 0.1
    var v34 int8 = 81
    if true {
        var v35 int64 = genFn1()
        v35 = (v3 % ((v35) | 1))
        _ = len(v6)
        _ = v35
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v4["a"] = genFn0()
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
