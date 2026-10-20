package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint16 }
func genFn0(p0 uint64, p1 uint, p2 bool) int32 {
    p2 = false
    for fi7 := 0; fi7 < 5; fi7++ {
        if (!(42 != (p0 << ((4294967295) & 63)))) {
            var v8 int64 = -803
            p1 = p1
            _ = v8
        }
        v9 := []int{}
        fmt.Println(p0)
        _ = fi7
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return -2147483648
}
func genFn1() int16 {
    var v10 string = (("hello" + (("a" + "\n\t") + ("a" + ("" + ("foo bar" + "a"))))) + ("\n\t" + "hello"))
    v10 = (v10 + v10)
    if (559 < 0) {
        for fi11 := 0; fi11 < 5; fi11++ {
            v10 = v10
            if false {
                v12 := []int{(^(fi11)), fi11, fi11, fi11}
                v10 = "foo bar"
                v10 = "\n\t"
                _ = v12
            }
            for fi13 := 0; fi13 < 4; fi13++ {
                v10 = ((v10 + v10) + v10)
                var v14 int8 = -128
                _ = fi13
                _ = v14
            }
            _ = fi11
        }
        v10 = v10
    }
    fmt.Println(v10)
    var v15 uint16 = 32767
    _ = v10
    _ = v15
    return 1000
}

func main() {
    var v1 uint32 = 42
    var v2 int16 = 868
    var v3 uint = 792
    var v4 string = ("foo bar" + ("hello" + "\n\t"))
    var v5 float64 = 1e10
    var v6 string = v4
    v16 := []int{4294967295, -325}
    for fi17 := 0; fi17 < 18; fi17++ {
        v4 = "a"
        v4 = "a"
        if false {
            v2 = (v2 + v2)
            v16 = append(v16, (fi17 % (((fi17 % (((fi17 + 2147483647)) | 1))) | 1)))
            var v18 bool = (3.14159 == 0.0)
            _ = v18
        } else {
            v19 := GenS1{F0: (945 > -508)}
            _ = v19
        }
        _ = fi17
    }
    var v20 int = -297
    var v21 uint32 = (v1 * 1)
    v1 = 254
    fmt.Println(v3, v5)
    switch v6 {
    case "foo bar":
        if (32767 <= 32767) {
            var v22 float64 = (((v5 / v5) / 3.14159) / float64(-2.25))
            for fi23 := 0; fi23 < 1; fi23++ {
                fmt.Println(v1, v22, v3)
                _ = len(v16)
                if false {
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    if false {
                        var v24 uint64 = 861
                        _ = v24
                    }
                    switch (v20 + 32768) {
                    case 2147483647:
                        var v25 uint = v3
                        _ = v25
                    case 127:
                        if len(v16) > 0 {
                            v16[0] = (fi23 << ((466) & 63))
                        }
                    case 32767:
                        switch ("a" + "hello") {
                        case "hello":
                            var v26 int32 = 26
                            _ = v26
                        case "\n\t":
                            if true {
                                for fi27 := 0; fi27 < 9; fi27++ {
                                    v6 = ""
                                    _ = fi27
                                }
                                var v28 uint16 = 1
                                v29 := GenS1{F0: false}
                                _ = v28
                                _ = v29
                            } else {
                                v30 := GenS1{F0: true}
                                _ = v30
                            }
                        case "foo bar":
                            v20 = (fi23 - fi23)
                        }
                    }
                }
                _ = fi23
            }
            _ = v22
        }
    case "\n\t":
        v20 = v20
    case "":
        if (true && (!false)) {
            v21 = uint32(v3)
        } else {
            v6 = ("hello" + v4)
            fmt.Println(v6)
        }
    default:
        v16 = append(v16, -583)
    }
    var v31 string = "hello"
    switch v20 {
    case 9223372036854775807:
        v32 := GenS2{F0: (v31 + "hello"), F1: 32767}
        _ = v32
    case -931:
        var v33 int = v20
        _ = v33
    case 803:
        fmt.Println(v31, v6)
    }
    v4 = (v4 + v6)
    fmt.Println(v4, v2, v4)
    v34 := GenS2{F0: v31, F1: 2}
    for fi35 := 0; fi35 < 7; fi35++ {
        v1 = v21
        _ = fi35
    }
    var v36 int32 = -414
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v31
    _ = v34
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
