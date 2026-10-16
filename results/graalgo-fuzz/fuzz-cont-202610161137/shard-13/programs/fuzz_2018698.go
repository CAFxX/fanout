package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint16 }
type GenS2 struct { F0 int8; F1 uint8; F2 bool }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 int16 }

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 2147483647
    var v3 int64 = -9223372036854775808
    v4 := []int{}
    v5 := []int{205}
    var v6 uint64 = 2
    fmt.Println(v2, v2)
    if false {
        var v7 bool = ((!((false || (-2.25 <= (((0.0 / 1.0) + (float64(0.1) * -2.25)) / 1.0))) && ((128 == -1) || (v6 <= v6)))) && false)
        _ = v7
    }
    delete(v1, "foo bar")
    v2 += v2
    v6 -= 0
    var v8 int = 32767
    var v9 int = (v8 >> ((35) & 63))
    switch ((("hello" + (("foo bar" + "a") + "a")) + "") + "hello") {
    case "foo bar":
        fmt.Println(v9, v8)
    default:
        var v10 int16 = 32767
        _ = v10
    }
    var v11 int16 = 32767
    var v12 bool = (!(v6 <= 4294967295))
    if len(v5) > 0 {
        v5[0] = 0
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        v12 = true
        var v14 int32 = int32(v6)
        _ = fi13
        _ = v14
    }
    var v15 int8 = -45
    for fi16 := 0; fi16 < 13; fi16++ {
        v1["hello"] = v9
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if true {
            v6 = (v6 | (v6 / (((v6 - v6)) | 1)))
            var v17 string = ("foo bar" + (("a" + ("" + ("\n\t" + ("\n\t" + "foo bar")))) + ("a" + "a")))
            if v12 {
                v18 := GenS1{F0: (v2 - (v2 + int32(v11))), F1: v11, F2: 32767}
                if false {
                    _ = len(v1)
                    v19 := GenS1{F0: -2147483648, F1: (v11 & (v11 ^ 1)), F2: 255}
                    _ = len(v4)
                    _ = v19
                } else {
                    switch v17 {
                    case "a":
                        var v20 float32 = (1e10 / 1.0)
                        _ = v20
                    default:
                        v11 = v11
                    }
                    if (1 != v15) {
                        var v21 uint = 42
                        v9 += fi16
                        _ = v21
                    }
                    fmt.Println(v3)
                }
                if ((v3 - v3) == v3) {
                    for fi22 := 0; fi22 < 3; fi22++ {
                        var v23 bool = v12
                        var v24 string = v17
                        var v25 float32 = 0.1
                        var v26 int = fi22
                        v2 = v2
                        _ = fi22
                        _ = v23
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                }
                _ = v18
            }
            _ = v17
        }
        _ = fi16
    }
    delete(v1, "a")
    if (!((float32(3.14159) * float32((3.14159 + 3.14159))) <= 0.1)) {
        for fi27 := 0; fi27 < 20; fi27++ {
            for fi28 := 0; fi28 < 16; fi28++ {
                if v12 {
                    v29 := GenS2{F0: 114, F1: 60, F2: (!(!(false && ("hello" > ("hello" + "\n\t")))))}
                    var v30 uint32 = 0
                    if len(v5) > 0 {
                        v5[0] = -9223372036854775808
                    }
                    _ = v29
                    _ = v30
                }
                fmt.Println(v12)
                v12 = true
                if v12 {
                    v31 := GenS1{F0: -63, F1: (-(v11)), F2: 197}
                    v32 := v1["foo bar"]
                    v11 += v11
                    _ = v31
                    _ = v32
                } else {
                    fmt.Println(v9)
                    if (!(((0 <= v11) || ((v2 ^ v2) >= 0)) || v12)) {
                        fmt.Println(v6)
                        var v33 int16 = -32768
                        v8 += 32767
                        var v34 int = fi28
                        var v35 uint64 = v6
                        v2 += v2
                        _ = v33
                        _ = v34
                        _ = v35
                    } else {
                        switch fi27 {
                        case 42:
                            v36 := GenS1{F0: (v2 - v2), F1: (v11 & int16(v15)), F2: 0}
                            _ = v36
                        case 756:
                            _ = len(v1)
                        case 32768:
                            v37 := GenS3{F0: (v11 / ((0) | 1)), F1: 2147483647}
                            _ = v37
                        }
                        v15 -= (^(v15))
                        fmt.Println(v12, v8, v11)
                    }
                }
                delete(v1, "")
                v6 = v6
                _ = fi28
            }
            _ = fi27
        }
        if (true || (-2147483648 > int(v11))) {
            v38 := GenS1{F0: (v2 >> ((65535) & 63)), F1: (v11 & v11), F2: 2}
            _ = v38
        } else {
            v39 := GenS4{F0: int16(v2)}
            v40 := GenS3{F0: v11, F1: (v2 ^ v2)}
            _ = v39
            _ = v40
        }
    }
    fmt.Println(v8)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
