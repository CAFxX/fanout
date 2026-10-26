package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint16, p1 uint8) uint32 {
    switch "\n\t" {
    case "foo bar":
        p0 = (p0 + uint16(p1))
    case "\n\t":
        fmt.Println(p0, p0)
    default:
        v5 := []int{398, -1}
        _ = v5
    }
    switch 9223372036854775807 {
    case 42:
        for fi6 := 0; fi6 < 11; fi6++ {
            var v7 string = "hello"
            v8 := []int{fi6, 32767, (^(fi6))}
            switch (fi6 % (((fi6 & fi6)) | 1)) {
            case 563:
                p1 += (p1 >> ((858) & 63))
            default:
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            }
            _ = fi6
            _ = v7
            _ = v8
        }
    default:
        v9 := GenS1{F0: 1}
        _ = v9
    }
    _ = p0
    _ = p1
    return 65535
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 42
    v3 := []string{"a"}
    v4 := make(map[uint32]bool)
    fmt.Println(v2)
    v10 := GenS2{F0: 522, F1: 0}
    var v11 int8 = -128
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 string = ("a" + "foo bar")
        v13 = (v13 + (v13 + v13))
        v14 := GenS2{F0: 59, F1: 128}
        _ = fi12
        _ = v13
        _ = v14
    }
    v11 -= (v11 % ((v11) | 1))
    fmt.Println(v2)
    fmt.Println(v11, v11)
    var v15 int16 = 284
    var v16 int8 = 100
    v17 := GenS1{F0: -2147483648}
    _ = len(v3)
    switch ("hello" + ((((("\n\t" + "hello") + ("\n\t" + (("hello" + "\n\t") + ("a" + ("hello" + "\n\t"))))) + "\n\t") + "") + (("a" + ("\n\t" + "\n\t")) + ("hello" + ("\n\t" + "a"))))) {
    case "a":
        var v18 int = 32767
        _ = v18
    default:
        if true {
            switch "hello" {
            case "\n\t":
                v2 += 42
            }
            var v19 float64 = 0.1
            v2 = v2
            _ = v19
        } else {
            var v20 string = "a"
            if len(v3) > 0 {
                v3[0] = (v20 + ("\n\t" + v20))
            }
            switch ((((v20 + "hello") + v20) + ("hello" + v20)) + (("a" + ("a" + (v20 + (v20 + (v20 + "foo bar"))))) + v20)) {
            case "foo bar":
                var v21 int16 = v15
                _ = v21
            case "\n\t":
                if (32767 != v15) {
                    if (!((1.5 - 1e10) > 3.14159)) {
                        v11 += 0
                        v22 := GenS3{F0: 865, F1: 0}
                        _ = v22
                    } else {
                        fmt.Println(v20)
                        for fi23 := 0; fi23 < 19; fi23++ {
                            v24 := GenS2{F0: 565, F1: 255}
                            v1["foo bar"] = false
                            _ = fi23
                            _ = v24
                        }
                    }
                    delete(v1, "foo bar")
                } else {
                    if (float64(1.5) != 0.1) {
                        v15 = (v15 * (v15 * 42))
                        var v25 int64 = 9223372036854775807
                        v2 = v2
                        _ = v25
                    } else {
                        var v26 uint16 = 761
                        fmt.Println(v16, v16, v15)
                        var v27 int8 = -39
                        _ = v26
                        _ = v27
                    }
                    v28 := GenS1{F0: 0}
                    _ = v28
                }
            }
            _ = v20
        }
    }
    v16 = 100
    for fi29 := 0; fi29 < 2; fi29++ {
        v2 = 0
        v30 := GenS3{F0: 0, F1: 2147483648}
        switch "\n\t" {
        case "":
            switch (fi29 >> ((926) & 63)) {
            case 1:
                var v31 uint = 255
                _ = v31
            case 127:
                var v32 int = 9223372036854775807
                _ = v32
            case 2147483647:
                v16 = 100
            default:
                v33 := GenS1{F0: 42}
                _ = v33
            }
        case "\n\t":
            var v34 uint8 = 84
            _ = v34
        default:
            v35 := GenS3{F0: 0, F1: genFn0(255, 42)}
            _ = v35
        }
        _ = fi29
        _ = v30
    }
    var v36 int16 = v15
    var v37 uint16 = 42
    for fi38 := 0; fi38 < 20; fi38++ {
        switch (fi38 % ((fi38) | 1)) {
        case 9223372036854775807:
            fmt.Println(v37)
        case 0:
            fmt.Println(v11, v16, fi38)
        case 32767:
            for fi39 := 0; fi39 < 1; fi39++ {
                var v40 uint8 = 127
                var v41 uint32 = 1
                _ = fi39
                _ = v40
                _ = v41
            }
        }
        _ = fi38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
