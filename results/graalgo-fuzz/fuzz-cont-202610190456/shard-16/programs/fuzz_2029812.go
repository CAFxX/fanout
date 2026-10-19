package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 uint; F2 uint16 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 bool; F1 int32; F2 bool }
func genFn0(p0 uint32, p1 uint16, p2 uint16) string {
    var v4 float64 = 0.1
    v5 := []int{439, -2147483648}
    switch 128 {
    case 1:
        if false {
            _ = len(v5)
        }
    case -1:
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return ("" + ("\n\t" + "hello"))
}
func genFn1(p0 uint8, p1 int, p2 uint64) int8 {
    if true {
        v6 := []int{p1}
        var v7 uint16 = 65535
        _ = v6
        _ = v7
    } else {
        v8 := []int{-1, 32767, p1, p1}
        _ = len(v8)
        _ = v8
    }
    v9 := []int{}
    fmt.Println(p1)
    var v10 uint16 = 2
    var v11 uint32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return int8(p1)
}

func main() {
    var v1 uint64 = 428
    v2 := []string{"a", "a", "a", ""}
    var v3 string = (("" + "a") + (("" + "a") + ("a" + "foo bar")))
    v3 = v3
    v12 := GenS1{F0: 647}
    v13 := GenS1{F0: 2}
    switch 669 {
    case 5:
        v2 = append(v2, "hello")
    default:
        fmt.Println(v3, v3)
    }
    v14 := GenS2{F0: 9223372036854775807, F1: 0, F2: 747}
    v15 := GenS3{F0: -32768, F1: 100}
    for fi16 := 0; fi16 < 6; fi16++ {
        var v17 uint = 713
        var v18 uint8 = uint8(v17)
        fmt.Println(fi16)
        _ = fi16
        _ = v17
        _ = v18
    }
    v3 = v3
    switch ((v3 + ("hello" + "")) + ((("a" + v3) + (v3 + (v3 + v3))) + v3)) {
    case "a":
        if false {
            fmt.Println(v1, v3)
            for fi19 := 0; fi19 < 11; fi19++ {
                if false {
                    v1 = v1
                    fmt.Println(v3, fi19)
                }
                fmt.Println(v1, v1)
                var v20 float64 = (3.14159 - 1.5)
                _ = fi19
                _ = v20
            }
            var v21 int8 = 42
            _ = v21
        } else {
            var v22 uint8 = 1
            var v23 int64 = 42
            switch -2147483648 {
            case 4294967295:
                for fi24 := 0; fi24 < 10; fi24++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi24
                }
            case 32767:
                var v25 int = 0
                _ = v25
            }
            _ = v22
            _ = v23
        }
    case "\n\t":
        fmt.Println(v1)
    case "foo bar":
        fmt.Println(v3)
    default:
        for fi26 := 0; fi26 < 1; fi26++ {
            switch (fi26 << ((18446744073709551615) & 63)) {
            case 32767:
                v27 := GenS1{F0: 1}
                _ = v27
            case -2147483648:
                for fi28 := 0; fi28 < 1; fi28++ {
                    if (1 >= 913) {
                        v3 = (("foo bar" + (v3 + (((v3 + ((("\n\t" + v3) + v3) + "foo bar")) + ((v3 + "a") + v3)) + v3))) + v3)
                        if true {
                            var v29 string = v3
                            _ = len(v2)
                            _ = v29
                        } else {
                            if (42 != -1) {
                                v2 = append(v2, "")
                                var v30 bool = (float32(3.14159) >= 3.14159)
                                _ = v30
                            }
                        }
                    }
                    _ = fi28
                }
            case -1:
                if true {
                    var v31 uint = 298
                    v32 := GenS2{F0: -435, F1: 960, F2: 65535}
                    _ = v31
                    _ = v32
                } else {
                    v1 = 168
                    fmt.Println(v3, v1)
                }
            default:
                for fi33 := 0; fi33 < 7; fi33++ {
                    switch 4294967295 {
                    case 9223372036854775807:
                        var v34 float32 = 1.5
                        _ = v34
                    case 2147483647:
                        v35 := GenS4{F0: 2, F1: (42 < -1)}
                        _ = v35
                    case 4294967295:
                        var v36 uint16 = 970
                        _ = v36
                    default:
                        var v37 int32 = 42
                        _ = v37
                    }
                    if false {
                        var v38 uint16 = 358
                        _ = v38
                    } else {
                        switch fi33 {
                        case 42:
                            v2 = append(v2, (v3 + "\n\t"))
                        default:
                            v1 = v1
                        }
                        v1 += 42
                        v39 := GenS5{F0: ((!true) || (42 != 42)), F1: 0, F2: true}
                        _ = v39
                    }
                    var v40 uint32 = 65535
                    _ = fi33
                    _ = v40
                }
            }
            v41 := GenS5{F0: false, F1: -1, F2: false}
            fmt.Println(fi26, v3, v1)
            _ = fi26
            _ = v41
        }
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
