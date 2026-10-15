package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }
func genFn0() uint16 {
    if true {
        switch (("foo bar" + ("a" + ("foo bar" + "a"))) + "a") {
        case "":
            var v6 int32 = 859
            _ = v6
        case "hello":
            var v7 float64 = 1.5
            _ = v7
        case "a":
            if false {
                var v8 uint = 42
                var v9 int16 = 32767
                v8 = 0
                _ = v8
                _ = v9
            } else {
                var v10 int = 42
                v11 := []int{(^(v10))}
                _ = v10
                _ = v11
            }
        }
        v12 := GenS1{F0: 4294967295, F1: 134}
        var v13 uint16 = 273
        _ = v12
        _ = v13
    }
    var v14 float32 = float32(3.14159)
    switch 334 {
    case -2147483648:
        v15 := GenS1{F0: 474, F1: 77}
        _ = v15
    case -10:
        for fi16 := 0; fi16 < 2; fi16++ {
            for fi17 := 0; fi17 < 11; fi17++ {
                if ((((1.5 == -2.25) && true) && (127 >= 100)) && true) {
                    if (757 >= 1) {
                        v18 := GenS1{F0: 2, F1: 181}
                        v19 := GenS1{F0: 18446744073709551615, F1: 2}
                        v14 = float32(-2.25)
                        v20 := GenS2{F0: 2}
                        _ = v18
                        _ = v19
                        _ = v20
                    } else {
                        fmt.Println(v14)
                        fmt.Println(fi16)
                    }
                    var v21 int64 = 410
                    _ = v21
                }
                v14 = 0.0
                _ = fi17
            }
            fmt.Println(v14, v14)
            if true {
                v22 := GenS1{F0: 2, F1: 128}
                v23 := GenS2{F0: 2}
                _ = v22
                _ = v23
            } else {
                var v24 int64 = 42
                var v25 int16 = int16(fi16)
                _ = v24
                _ = v25
            }
            _ = fi16
        }
    case 42:
        fmt.Println(v14)
    }
    _ = v14
    return 521
}
func genFn1(p0 uint32) bool {
    var v26 uint8 = 1
    p0 += p0
    _ = p0
    _ = v26
    return (349 <= 345)
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 9223372036854775807
    var v3 int32 = 0
    var v4 int32 = v3
    var v5 uint16 = 42
    v27 := []int{128, 4294967295}
    if len(v27) >= 2 {
        v27 = v27[1:len(v27)]
    }
    for fi28 := 0; fi28 < 15; fi28++ {
        v29 := GenS3{F0: "\n\t"}
        for fi30 := 0; fi30 < 16; fi30++ {
            delete(v1, fi28)
            _ = len(v27)
            var v31 float64 = -2.25
            _ = fi30
            _ = v31
        }
        _ = fi28
        _ = v29
    }
    fmt.Println(v4)
    fmt.Println(v5)
    fmt.Println(v3, v2, v2)
    var v32 float64 = 1e10
    for fi33 := 0; fi33 < 1; fi33++ {
        var v34 uint = 18446744073709551615
        if true {
            v2 += v2
        } else {
            for fi35 := 0; fi35 < 3; fi35++ {
                if (true && false) {
                    v4 += v4
                    v4 -= 152
                    var v36 uint = 0
                    v5 = genFn0()
                    _ = v36
                }
                var v37 bool = true
                var v38 float64 = 1e10
                _ = len(v27)
                var v39 string = ("foo bar" + "a")
                _ = fi35
                _ = v37
                _ = v38
                _ = v39
            }
            fmt.Println(fi33, v32)
        }
        _ = fi33
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
