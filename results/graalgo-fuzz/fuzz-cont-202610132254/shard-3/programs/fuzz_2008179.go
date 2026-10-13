package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint }
type GenS2 struct { F0 int; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint64; F1 uint }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]bool)
    var v3 uint = 1
    v4 := make(map[string]string)
    var v5 int32 = -791
    v5 = int32(v3)
    switch (("hello" + "\n\t") + "hello") {
    case "foo bar":
        fmt.Println(v5)
    case "\n\t":
        _ = len(v2)
    }
    fmt.Println(v5)
    v5 += 42
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            v8 := []int{}
            switch ((("" + "a") + "a") + "foo bar") {
            case "":
                var v9 float32 = (3.14159 - (float32(float64(0.1)) - (float32(0.1) / 0.1)))
                _ = v9
            case "a":
                v3 += v3
            case "foo bar":
                var v10 uint = v3
                _ = v10
            }
            _ = fi7
            _ = v8
        }
        if false {
            v5 = (-(v5))
            var v11 float64 = float64(-2.25)
            v3 += (v3 % (((v3 & v3)) | 1))
            _ = v11
        } else {
            v3 = (^(v3))
            if true {
                for fi12 := 0; fi12 < 7; fi12++ {
                    v13 := GenS1{F0: 704, F1: 1, F2: v3}
                    _ = fi12
                    _ = v13
                }
            } else {
                v5 = (v5 + v5)
                var v14 string = ("hello" + "\n\t")
                if ((-2.25 + ((-2.25 - 0.1) + (3.14159 + ((0.1 / 1.0) * float64(float32(0.0)))))) > 1e10) {
                    if (!true) {
                        switch (-(fi6)) {
                        case 9223372036854775807:
                            var v15 float64 = -2.25
                            _ = v15
                        case -1:
                            var v16 bool = (((v14 + v14) >= "hello") || true)
                            _ = v16
                        default:
                            for fi17 := 0; fi17 < 7; fi17++ {
                                var v18 uint64 = 0
                                fmt.Println(v14)
                                _ = fi17
                                _ = v18
                            }
                        }
                        v3 += v3
                    } else {
                        if true {
                            delete(v2, (fi6 >> ((v3) & 63)))
                        }
                    }
                    if true {
                        if true {
                            var v19 float64 = (-2.25 + 1e10)
                            v20 := []int{fi6}
                            _ = v19
                            _ = v20
                        }
                    } else {
                        v14 = v14
                        for fi21 := 0; fi21 < 6; fi21++ {
                            v5 += v5
                            var v22 uint32 = 522
                            _ = fi21
                            _ = v22
                        }
                    }
                } else {
                    v3 = (v3 + v3)
                    if (false || (int16(v5) <= -985)) {
                        v23 := []int{377, fi6, (-(fi6))}
                        _ = v23
                    }
                    v24 := []int{(fi6 << ((v3) & 63)), fi6, 32767}
                    _ = v24
                }
                _ = v14
            }
        }
        _ = fi6
    }
    var v25 int16 = int16(v3)
    v26 := GenS1{F0: 1, F1: -9223372036854775808, F2: v3}
    for fi27 := 0; fi27 < 8; fi27++ {
        fmt.Println(v25)
        v28 := GenS2{F0: 127, F1: 42, F2: -1}
        delete(v2, fi27)
        _ = fi27
        _ = v28
    }
    v25 = (v25 << ((65535) & 63))
    v29 := GenS3{F0: 0, F1: v3}
    var v30 uint32 = 2147483647
    v25 += (v25 >> ((v3) & 63))
    if true {
        var v31 uint32 = 523
        for fi32 := 0; fi32 < 1; fi32++ {
            for fi33 := 0; fi33 < 4; fi33++ {
                fmt.Println(v31)
                _ = fi33
            }
            fmt.Println(v3)
            _ = fi32
        }
        v25 = -1
        _ = v31
    } else {
        v34 := GenS1{F0: 65535, F1: 42, F2: (v3 / ((42) | 1))}
        _ = v34
    }
    var v35 uint64 = 4294967295
    if (false && false) {
        v35 = 2
        v25 -= (v25 << ((279) & 63))
    } else {
        var v36 uint8 = 202
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v4
    _ = v5
}
