package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int }
type GenS2 struct { F0 uint16; F1 int32 }

func main() {
    var v1 string = "hello"
    var v2 int64 = -193
    v3 := make(map[uint32]uint32)
    var v4 int16 = 42
    var v5 uint8 = 0
    for fi6 := 0; fi6 < 16; fi6++ {
        v5 = 63
        _ = len(v3)
        _ = fi6
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 uint16 = 401
        v1 = (v1 + (v1 + (v1 + (("foo bar" + v1) + ""))))
        var v9 int16 = 32767
        _ = fi7
        _ = v8
        _ = v9
    }
    switch ((v1 + v1) + v1) {
    case "hello":
        for fi10 := 0; fi10 < 2; fi10++ {
            for fi11 := 0; fi11 < 6; fi11++ {
                switch fi10 {
                case -1:
                    v4 = v4
                case 128:
                    v12 := GenS1{F0: 341, F1: 80, F2: fi10}
                    _ = v12
                default:
                    var v13 uint16 = 0
                    _ = v13
                }
                v5 = (v5 - v5)
                if true {
                    _ = len(v3)
                    v1 = "foo bar"
                    if true {
                        fmt.Println(v2)
                        delete(v3, 2)
                    }
                } else {
                    fmt.Println(fi10)
                }
                _ = fi11
            }
            v2 -= (v2 ^ 180)
            v14 := []int{0, -1, -550}
            _ = fi10
            _ = v14
        }
    }
    var v15 string = v1
    for fi16 := 0; fi16 < 3; fi16++ {
        v4 = 0
        if (true || (42 < 42)) {
            v15 = ((v15 + ("hello" + ((v15 + v1) + ((("" + ("a" + (v1 + "foo bar"))) + (("hello" + v15) + "")) + v1)))) + "foo bar")
        } else {
            switch (fi16 - fi16) {
            case -540:
                for fi17 := 0; fi17 < 14; fi17++ {
                    var v18 uint16 = 0
                    _ = fi17
                    _ = v18
                }
            case -921:
                if ((int(v4) != -737) && ((v5 / ((v5) | 1)) <= v5)) {
                    for fi19 := 0; fi19 < 15; fi19++ {
                        v15 = (v15 + (v1 + v1))
                        var v20 int64 = v2
                        _ = fi19
                        _ = v20
                    }
                    var v21 string = ""
                    var v22 int32 = 776
                    _ = v21
                    _ = v22
                }
            case 32768:
                if (11 == -102) {
                    for fi23 := 0; fi23 < 7; fi23++ {
                        v24 := GenS1{F0: 110, F1: 2, F2: (fi23 * (-(fi23)))}
                        var v25 int = (fi23 / (((^(fi23))) | 1))
                        v15 = (v15 + v1)
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    fmt.Println(v5, fi16, v5)
                }
            }
            switch v1 {
            case "foo bar":
                v4 -= (v4 / (((v4 % (((v4 >> ((255) & 63))) | 1))) | 1))
            case "hello":
                v15 = v1
            }
        }
        var v26 uint64 = 710
        _ = fi16
        _ = v26
    }
    v15 = (v1 + "hello")
    if true {
        if false {
            v27 := GenS2{F0: 42, F1: -1}
            _ = v27
        }
        var v28 uint32 = 4294967295
        v29 := GenS2{F0: 2, F1: 1}
        _ = v28
        _ = v29
    } else {
        v5 = 42
        var v30 int16 = v4
        switch -205 {
        case 38:
            v31 := []int{1, 127, 4294967295, 888}
            _ = v31
        case -9223372036854775808:
            v3[635] = 343
        default:
            var v32 int16 = (v30 / (((-(v30))) | 1))
            _ = v32
        }
        _ = v30
    }
    for fi33 := 0; fi33 < 11; fi33++ {
        var v34 uint32 = 65535
        switch fi33 {
        case 127:
            fmt.Println(v15, v4)
        case -95:
            v35 := []int{}
            _ = v35
        default:
            v15 = v15
        }
        _ = fi33
        _ = v34
    }
    var v36 int = 32767
    v37 := []int{v36, (v36 ^ (v36 * (v36 / (((v36 * 852)) | 1)))), v36, 32767}
    if len(v37) > 0 {
        v37[0] = (v36 - (v36 << ((825) & 63)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
