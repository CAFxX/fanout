package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 bool; F1 uint8 }
func genFn0() int {
    v5 := []int{867, -286, 0, 127}
    if true {
        for fi6 := 0; fi6 < 16; fi6++ {
            v7 := GenS1{F0: 2147483647}
            v8 := GenS1{F0: 2}
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        if false {
            var v10 string = ("a" + ("" + ""))
            if false {
                if false {
                    if len(v5) > 0 {
                        v5[0] = 655
                    }
                }
                if false {
                    var v11 bool = (false && (true || (127 >= 136)))
                    for fi12 := 0; fi12 < 20; fi12++ {
                        if v11 {
                            _ = len(v5)
                            var v13 int32 = -1
                            _ = v13
                        }
                        _ = fi12
                    }
                    _ = v11
                } else {
                    fmt.Println(v10, fi9)
                    fmt.Println(v10, v10)
                    var v14 int8 = 1
                    _ = v14
                }
            } else {
                for fi15 := 0; fi15 < 20; fi15++ {
                    var v16 float32 = 0.1
                    _ = fi15
                    _ = v16
                }
                v10 = ("" + v10)
            }
            _ = v10
        }
        var v17 int64 = -9223372036854775808
        _ = fi9
        _ = v17
    }
    _ = v5
    return 9223372036854775807
}

func main() {
    v1 := []uint32{0}
    v2 := []int{}
    var v3 string = "foo bar"
    var v4 bool = true
    if false {
        for fi18 := 0; fi18 < 9; fi18++ {
            fmt.Println(v4, v4)
            if ((v3 <= v3) || false) {
                if ((float64(((1.5 + ((0.0 - 0.1) + 0.0)) * 1e10)) == 1.5) || v4) {
                    v19 := GenS1{F0: 544}
                    v20 := GenS2{F0: 1, F1: v3}
                    _ = v19
                    _ = v20
                } else {
                    v21 := GenS2{F0: 1, F1: ""}
                    _ = v21
                }
                var v22 int16 = -290
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v22
            } else {
                var v23 uint32 = 65535
                if v4 {
                    var v24 uint8 = 0
                    _ = v24
                }
                _ = v23
            }
            fmt.Println(v3, v3)
            _ = fi18
        }
        var v25 bool = v4
        var v26 float32 = (float32(float64(0.0)) * 1.5)
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 3; fi27++ {
        fmt.Println(v3, fi27)
        fmt.Println(v3, v3, v4)
        v28 := GenS3{F0: (!v4), F1: 2}
        _ = fi27
        _ = v28
    }
    _ = len(v2)
    v4 = (42 < 924)
    fmt.Println(v4)
    for fi29 := 0; fi29 < 6; fi29++ {
        for fi30 := 0; fi30 < 19; fi30++ {
            v4 = v4
            _ = len(v1)
            _ = fi30
        }
        _ = len(v2)
        _ = fi29
    }
    v31 := GenS3{F0: v4, F1: 128}
    fmt.Println(v3, v3)
    fmt.Println(v3)
    var v32 uint32 = 669
    var v33 float64 = 1.5
    var v34 float64 = v33
    var v35 uint8 = 223
    var v36 int16 = 32767
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
