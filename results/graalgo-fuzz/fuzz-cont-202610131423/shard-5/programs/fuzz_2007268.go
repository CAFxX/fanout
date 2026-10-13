package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string }
func genFn0(p0 uint64) int32 {
    fmt.Println(p0)
    p0 += 18446744073709551615
    var v4 uint = 4294967295
    if false {
        var v5 uint32 = 0
        v4 = v4
        fmt.Println(v5)
        _ = v5
    } else {
        v6 := []int{-729, 32767, 0}
        var v7 uint8 = 126
        _ = v6
        _ = v7
    }
    p0 -= (p0 ^ 0)
    _ = p0
    _ = v4
    return 1
}
func genFn1(p0 int8, p1 string) uint64 {
    fmt.Println(p1)
    v8 := GenS1{F0: -1}
    p0 += 1
    _ = p0
    _ = p1
    _ = v8
    return 4294967295
}

func main() {
    v1 := []uint32{}
    v2 := make(map[string]int)
    var v3 string = ("hello" + "hello")
    switch v3 {
    case "a":
        v3 = v3
    case "":
        var v9 uint8 = 1
        _ = v9
    case "foo bar":
        if true {
            if (42 != 1) {
                v3 = v3
                for fi10 := 0; fi10 < 17; fi10++ {
                    var v11 uint64 = 798
                    _ = fi10
                    _ = v11
                }
            }
            var v12 int32 = -869
            v1 = append(v1, 0)
            _ = v12
        } else {
            for fi13 := 0; fi13 < 4; fi13++ {
                v14 := GenS2{F0: 4294967295}
                _ = fi13
                _ = v14
            }
            v15 := v2[("a" + ((v3 + v3) + "a"))]
            _ = v15
        }
    }
    if false {
        if (1000 != 968) {
            if false {
                v3 = "hello"
                fmt.Println(v3)
            } else {
                v2[v3] = -9223372036854775808
                var v16 int16 = 0
                _ = v16
            }
        }
        for fi17 := 0; fi17 < 19; fi17++ {
            v18 := GenS3{F0: -1}
            var v19 int64 = -590
            _ = fi17
            _ = v18
            _ = v19
        }
        for fi20 := 0; fi20 < 18; fi20++ {
            v3 = ""
            _ = fi20
        }
    } else {
        v21 := GenS4{F0: v3}
        v22 := GenS2{F0: 0}
        v3 = "a"
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        fmt.Println(fi23)
        _ = fi23
    }
    var v24 int8 = 74
    for fi25 := 0; fi25 < 19; fi25++ {
        v3 = (((((v3 + (v3 + "foo bar")) + v3) + (v3 + v3)) + ("" + v3)) + v3)
        _ = fi25
    }
    v1 = append(v1, 198)
    var v26 float32 = (1e10 * ((3.14159 + (float32((1.5 * 3.14159)) + 1.5)) * 0.1))
    switch -1 {
    case 9223372036854775807:
        v27 := GenS1{F0: -32768}
        _ = v27
    case 4294967295:
        v3 = (("hello" + "") + v3)
    }
    switch "foo bar" {
    case "":
        v2["a"] = -150
    case "foo bar":
        v24 -= (v24 >> ((1) & 63))
    }
    v28 := v2["a"]
    if false {
        switch v28 {
        case 9223372036854775807:
            if (genFn0(828) <= 2147483647) {
                var v29 uint8 = 127
                fmt.Println(v3, v29)
                _ = v29
            }
        case 553:
            if false {
                for fi30 := 0; fi30 < 3; fi30++ {
                    var v31 int16 = 1000
                    _ = fi30
                    _ = v31
                }
                if true {
                    if true {
                        _ = len(v1)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        v28 = v28
                    }
                    fmt.Println(v28, v3)
                } else {
                    v28 += v28
                }
                v32 := GenS1{F0: 0}
                _ = v32
            } else {
                v3 = ""
            }
        case 4294967295:
            switch v3 {
            case "":
                v24 = v24
            }
        }
    } else {
        switch -293 {
        case 32768:
            if (4294967295 >= 259) {
                var v33 uint16 = 664
                _ = len(v2)
                _ = v33
            }
        case 634:
            var v34 int64 = -938
            _ = v34
        case -9223372036854775808:
            var v35 int32 = 759
            _ = v35
        }
    }
    var v36 float32 = 1.5
    var v37 uint = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v28
    _ = v3
    _ = v36
    _ = v37
}
