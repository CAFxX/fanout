package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 string; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int64 }
func genFn0(p0 int16, p1 int8, p2 uint16) int32 {
    var v5 uint16 = 255
    v6 := GenS1{F0: 2147483647}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -2147483648
}

func main() {
    v1 := []int{}
    var v2 bool = (0 == 525)
    v3 := make(map[uint32]string)
    var v4 bool = (v2 && ((("a" <= ("foo bar" + "foo bar")) && (1 < 549)) && (v2 || true)))
    var v7 int32 = 42
    v8 := v3[2147483648]
    for fi9 := 0; fi9 < 2; fi9++ {
        if false {
            v10 := GenS2{F0: 0, F1: (v8 + v8), F2: 831}
            _ = v10
        } else {
            var v11 string = v8
            _ = v11
        }
        switch v8 {
        case "foo bar":
            v12 := GenS3{F0: 86, F1: 16}
            _ = v12
        default:
            if v4 {
                v13 := GenS1{F0: v7}
                fmt.Println(v2, fi9, v7)
                _ = v13
            } else {
                var v14 float64 = 0.0
                v15 := GenS2{F0: 1000, F1: ("hello" + v8), F2: -764}
                _ = v14
                _ = v15
            }
        }
        for fi16 := 0; fi16 < 5; fi16++ {
            fmt.Println(v4, v7, v2)
            v17 := GenS4{F0: -9223372036854775808}
            _ = fi16
            _ = v17
        }
        _ = fi9
    }
    switch ((v8 + "\n\t") + "a") {
    case "hello":
        fmt.Println(v8)
    case "foo bar":
        v18 := GenS4{F0: 0}
        _ = v18
    case "a":
        if (55 != 4294967295) {
            for fi19 := 0; fi19 < 19; fi19++ {
                if (v4 || false) {
                    if (!v2) {
                        _ = len(v3)
                    } else {
                        v20 := v3[2]
                        _ = v20
                    }
                    v21 := GenS3{F0: 214, F1: 0}
                    _ = v21
                }
                var v22 uint = 18446744073709551615
                _ = fi19
                _ = v22
            }
        } else {
            var v23 string = v8
            if v2 {
                if (-1 < -1) {
                    var v24 int32 = v7
                    var v25 int16 = 42
                    _ = v24
                    _ = v25
                }
                if v4 {
                    v26 := GenS1{F0: v7}
                    if (true && (false || (0 != 42))) {
                        var v27 int32 = 1
                        v2 = (!(908 <= 0))
                        _ = len(v3)
                        _ = v27
                    } else {
                        var v28 bool = (1 == 42)
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 1; fi29++ {
                        var v30 int8 = 0
                        fmt.Println(v7, fi29, v4)
                        switch (("hello" + "hello") + v8) {
                        case "foo bar":
                            v1 = append(v1, 101)
                        }
                        _ = fi29
                        _ = v30
                    }
                    _ = v26
                }
            }
            var v31 uint64 = 4294967295
            _ = v23
            _ = v31
        }
    }
    v32 := GenS1{F0: v7}
    var v33 int8 = 1
    var v34 uint = 4294967295
    var v35 uint = v34
    v3[4294967295] = (v8 + v8)
    var v36 uint64 = 4294967295
    for fi37 := 0; fi37 < 3; fi37++ {
        v4 = v4
        v38 := GenS1{F0: 2147483647}
        _ = fi37
        _ = v38
    }
    var v39 bool = (true || true)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v7
    _ = v8
}
