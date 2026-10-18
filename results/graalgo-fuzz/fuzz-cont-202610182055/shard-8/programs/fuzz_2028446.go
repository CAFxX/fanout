package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 int8; F2 int16 }
type GenS3 struct { F0 int64; F1 uint8; F2 int16 }
func genFn0(p0 uint, p1 bool, p2 int16) uint64 {
    switch ("a" + ((("a" + "foo bar") + "hello") + "")) {
    case "hello":
        var v6 uint64 = 18446744073709551615
        _ = v6
    }
    p0 = (p0 << ((p0) & 63))
    var v7 int32 = 42
    if p1 {
        p0 += 2
        var v8 int64 = -1
        _ = v8
    } else {
        switch "foo bar" {
        case "a":
            fmt.Println(p1)
        case "hello":
            for fi9 := 0; fi9 < 1; fi9++ {
                var v10 float32 = -2.25
                if p1 {
                    v10 = 0.1
                    v11 := []int{fi9, fi9}
                    var v12 int = (fi9 - 128)
                    _ = v11
                    _ = v12
                } else {
                    for fi13 := 0; fi13 < 13; fi13++ {
                        var v14 string = ("a" + "foo bar")
                        v15 := []int{(fi9 / (((fi9 >> (((p0 >> (((p0 + 260)) & 63))) & 63))) | 1)), fi9}
                        v16 := GenS1{F0: -106}
                        _ = fi13
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                    if true {
                        var v17 uint32 = 373
                        p1 = p1
                        if (4294967295 > 42) {
                            var v18 int64 = -9223372036854775808
                            if (float32(-2.25) >= (v10 - (v10 / v10))) {
                                v19 := GenS1{F0: 1}
                                switch (((("foo bar" + "") + "hello") + ("hello" + (((("" + (("foo bar" + ("hello" + "foo bar")) + "\n\t")) + ("foo bar" + "")) + "\n\t") + ""))) + ("a" + "a")) {
                                case "\n\t":
                                    var v20 string = "a"
                                    _ = v20
                                }
                                for fi21 := 0; fi21 < 9; fi21++ {
                                    fmt.Println(v10, p0)
                                    _ = fi21
                                }
                                _ = v19
                            }
                            _ = v18
                        } else {
                            v10 -= (v10 / 3.14159)
                            for fi22 := 0; fi22 < 12; fi22++ {
                                var v23 string = "hello"
                                v7 = v7
                                _ = fi22
                                _ = v23
                            }
                            var v24 int16 = 325
                            _ = v24
                        }
                        _ = v17
                    }
                    p1 = (!false)
                }
                _ = fi9
                _ = v10
            }
        }
        if true {
            p2 = (p2 / (((p2 & p2)) | 1))
        } else {
            var v25 int64 = 42
            v26 := []int{}
            _ = v25
            _ = v26
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 4294967295
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]string)
    var v3 uint16 = 1
    var v4 bool = false
    v5 := []uint32{65535, 4294967295}
    var v27 int = 1
    v28 := GenS1{F0: -16}
    var v29 float32 = -2.25
    for fi30 := 0; fi30 < 5; fi30++ {
        v31 := GenS2{F0: 1, F1: 118, F2: 859}
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 5; fi32++ {
        var v33 int = (v27 % (((fi32 - (v27 * v27))) | 1))
        _ = fi32
        _ = v33
    }
    v4 = (true && true)
    if len(v1) > 0 {
        v1[0] = 65535
    }
    var v34 uint8 = 2
    v35 := GenS3{F0: -903, F1: (v34 << ((42) & 63)), F2: -345}
    var v36 uint8 = v34
    var v37 uint = 65535
    fmt.Println(v3)
    v27 = v27
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
