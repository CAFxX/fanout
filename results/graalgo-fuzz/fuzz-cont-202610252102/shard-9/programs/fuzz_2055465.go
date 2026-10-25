package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 bool = (!false)
    var v2 float64 = -2.25
    v3 := make(map[string]bool)
    var v4 uint32 = 1
    v5 := []string{"hello"}
    fmt.Println(v4)
    fmt.Println(v2, v2)
    v1 = (1000 >= 1000)
    v3[""] = (!(3.14159 > 1e10))
    v4 -= v4
    v6 := GenS1{F0: 0, F1: 1}
    var v7 float64 = float64((1e10 + 0.0))
    v8 := GenS1{F0: 42, F1: 777}
    for fi9 := 0; fi9 < 7; fi9++ {
        v4 = (v4 - v4)
        _ = fi9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 int16 = -32768
        var v12 int32 = -2147483648
        _ = fi10
        _ = v11
        _ = v12
    }
    if (9223372036854775807 < 0) {
        v2 += v7
        if (458 != 42) {
            fmt.Println(v2, v7, v4)
            for fi13 := 0; fi13 < 7; fi13++ {
                if (42 == 100) {
                    if ((v7 / v2) == 1.5) {
                        var v14 uint64 = 0
                        for fi15 := 0; fi15 < 11; fi15++ {
                            var v16 string = (("foo bar" + ("hello" + (("" + ("foo bar" + (("a" + (("" + "hello") + (("foo bar" + "hello") + "hello"))) + "hello"))) + (("hello" + "a") + "hello")))) + "a")
                            var v17 int32 = -1
                            if (1 == 1) {
                                fmt.Println(v1, v1, fi13)
                            }
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                        _ = v14
                    }
                    if v1 {
                        fmt.Println(v1)
                        v18 := v3["hello"]
                        _ = v18
                    }
                    var v19 int = (fi13 | 42)
                    _ = v19
                }
                _ = fi13
            }
            switch -309 {
            case 127:
                v1 = (14 == v4)
            default:
                if v1 {
                    delete(v3, "a")
                    v20 := GenS1{F0: 539, F1: 1}
                    if v1 {
                        for fi21 := 0; fi21 < 18; fi21++ {
                            for fi22 := 0; fi22 < 4; fi22++ {
                                v23 := GenS2{F0: 1, F1: 75, F2: -13}
                                _ = len(v3)
                                v24 := GenS1{F0: 18446744073709551615, F1: 42}
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                            for fi25 := 0; fi25 < 12; fi25++ {
                                v2 = (v2 - ((v7 + 1e10) - (v2 + float64(1e10))))
                                _ = fi25
                            }
                            _ = fi21
                        }
                    }
                    _ = v20
                } else {
                    if v1 {
                        var v26 float32 = float32(float64(0.1))
                        for fi27 := 0; fi27 < 13; fi27++ {
                            v28 := GenS1{F0: 472, F1: 4294967295}
                            var v29 uint8 = 42
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                        for fi30 := 0; fi30 < 2; fi30++ {
                            v31 := GenS2{F0: 42, F1: 2, F2: (fi30 | fi30)}
                            v4 = 2
                            _ = fi30
                            _ = v31
                        }
                        _ = v26
                    } else {
                        for fi32 := 0; fi32 < 12; fi32++ {
                            if false {
                                fmt.Println(v2)
                                for fi33 := 0; fi33 < 20; fi33++ {
                                    var v34 uint32 = 42
                                    switch (("hello" + "a") + "hello") {
                                    case "":
                                        _ = len(v3)
                                    case "foo bar":
                                        v35 := GenS2{F0: 0, F1: 2, F2: fi33}
                                        _ = v35
                                    case "\n\t":
                                        switch "\n\t" {
                                        case "":
                                            var v36 uint16 = 0
                                            _ = v36
                                        default:
                                            var v37 uint = 2
                                            var v38 int32 = 42
                                            fmt.Println(v37, fi32, v7)
                                            _ = v37
                                            _ = v38
                                        }
                                    }
                                    _ = fi33
                                    _ = v34
                                }
                                for fi39 := 0; fi39 < 3; fi39++ {
                                    var v40 string = "hello"
                                    _ = fi39
                                    _ = v40
                                }
                            }
                            _ = fi32
                        }
                    }
                    v7 = (v7 / float64((0.0 + float32(-2.25))))
                }
            }
        } else {
            switch 9223372036854775807 {
            case 182:
                switch 115 {
                case 9223372036854775807:
                    var v41 float32 = -2.25
                    _ = v41
                }
            case 4294967295:
                fmt.Println(v2, v7, v7)
            }
            var v42 int16 = -1
            switch ("" + "\n\t") {
            case "a":
                v42 -= v42
            case "hello":
                v4 = (^(v4))
            case "":
                for fi43 := 0; fi43 < 10; fi43++ {
                    v7 -= (0.1 * -2.25)
                    fmt.Println(v1, v42, v4)
                    _ = fi43
                }
            }
            _ = v42
        }
        v44 := GenS3{F0: 65535}
        _ = v44
    } else {
        v45 := v3[(("a" + "foo bar") + "foo bar")]
        fmt.Println(v4, v1)
        _ = v45
    }
    var v46 int64 = 0
    _ = len(v3)
    for fi47 := 0; fi47 < 1; fi47++ {
        var v48 uint16 = 2
        v46 = -138
        v49 := GenS3{F0: v4}
        _ = fi47
        _ = v48
        _ = v49
    }
    var v50 uint64 = 18446744073709551615
    v4 += v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v50
    _ = v6
    _ = v7
    _ = v8
}
