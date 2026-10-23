package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint }
type GenS2 struct { F0 uint64; F1 int32; F2 int }
type GenS3 struct { F0 uint64 }
func genFn0() uint {
    var v6 uint64 = 1
    if false {
        var v7 uint = 4294967295
        v8 := GenS1{F0: 26, F1: 1, F2: v7}
        var v9 string = ("foo bar" + "")
        _ = v7
        _ = v8
        _ = v9
    }
    v6 -= (v6 / ((0) | 1))
    v10 := []int{}
    var v11 int64 = -1
    _ = v10
    _ = v11
    _ = v6
    return 18446744073709551615
}
func genFn1() uint64 {
    if false {
        v12 := []int{-173, 22, 32767}
        v12 = append(v12, 4294967295)
        _ = v12
    } else {
        if true {
            for fi13 := 0; fi13 < 5; fi13++ {
                v14 := []int{}
                _ = fi13
                _ = v14
            }
            for fi15 := 0; fi15 < 9; fi15++ {
                if (!((!(25 >= 2147483648)) && true)) {
                    if (false || (-128 != 127)) {
                        for fi16 := 0; fi16 < 6; fi16++ {
                            switch ("" + "") {
                            case "hello":
                                var v17 uint64 = 42
                                _ = v17
                            }
                            var v18 uint32 = 0
                            _ = fi16
                            _ = v18
                        }
                        switch ("\n\t" + "a") {
                        case "\n\t":
                            var v19 uint = 18446744073709551615
                            _ = v19
                        }
                    }
                    var v20 uint = 2
                    _ = v20
                }
                var v21 int = (fi15 << ((genFn0()) & 63))
                var v22 uint = 18446744073709551615
                _ = fi15
                _ = v21
                _ = v22
            }
        }
    }
    var v23 string = (("hello" + "") + ((("foo bar" + "hello") + "foo bar") + (("foo bar" + "a") + "hello")))
    var v24 uint = 4294967295
    _ = v23
    _ = v24
    return 18446744073709551615
}

func main() {
    var v1 int64 = 347
    v2 := make(map[int]string)
    var v3 float64 = -2.25
    v4 := []string{"foo bar", (((((("\n\t" + "a") + ("a" + ("foo bar" + "\n\t"))) + ((("a" + "foo bar") + "") + ((("" + "\n\t") + "a") + "\n\t"))) + "a") + ("" + "hello")) + "\n\t"), "\n\t"}
    v5 := []uint32{2, 229, 1}
    fmt.Println(v1, v3)
    v1 += v1
    v25 := GenS2{F0: 2, F1: 882, F2: -1}
    v26 := GenS1{F0: 2, F1: 2147483647, F2: 1}
    for fi27 := 0; fi27 < 7; fi27++ {
        if false {
            v1 += (^(v1))
            for fi28 := 0; fi28 < 2; fi28++ {
                if len(v5) > 0 {
                    v5[0] = 858
                }
                v1 -= v1
                if (true && (65535 == 426)) {
                    var v29 uint16 = 993
                    var v30 uint = 4294967295
                    var v31 int64 = 42
                    _ = v29
                    _ = v30
                    _ = v31
                } else {
                    switch fi28 {
                    case 127:
                        fmt.Println(v1)
                    case 228:
                        var v32 uint16 = 42
                        _ = v32
                    default:
                        v3 = (v3 + v3)
                    }
                    v3 = v3
                }
                _ = fi28
            }
        }
        var v33 float32 = 3.14159
        _ = fi27
        _ = v33
    }
    v34 := GenS3{F0: 42}
    v3 = v3
    if (float64(0.0) == v3) {
        v3 = 0.0
        v35 := GenS3{F0: genFn1()}
        switch "" {
        case "\n\t":
            for fi36 := 0; fi36 < 16; fi36++ {
                if (0.0 >= 0.0) {
                    v3 = v3
                    for fi37 := 0; fi37 < 9; fi37++ {
                        v3 = 0.0
                        _ = fi37
                    }
                } else {
                    switch (fi36 << ((249) & 63)) {
                    case 915:
                        v1 = (^(v1))
                    }
                    fmt.Println(v1, v3)
                    var v38 int64 = v1
                    _ = v38
                }
                _ = fi36
            }
        }
        _ = v35
    }
    if (-1 <= 127) {
        for fi39 := 0; fi39 < 11; fi39++ {
            var v40 uint = 4294967295
            fmt.Println(v3, v1)
            v41 := GenS1{F0: 128, F1: -760, F2: (v40 + 42)}
            _ = fi39
            _ = v40
            _ = v41
        }
        var v42 float64 = 1.5
        v2[32768] = ("foo bar" + "a")
        _ = v42
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
