package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
func genFn0() int64 {
    v7 := GenS1{F0: 255}
    var v8 int32 = -1
    _ = v7
    _ = v8
    return 0
}
func genFn1() uint {
    var v9 int = 32768
    var v10 int32 = -259
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 uint32 = 42
        v13 := []int{(v9 ^ fi11), -769}
        _ = fi11
        _ = v12
        _ = v13
    }
    v9 = 127
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v10, v10)
    case "hello":
        if (2 >= 2147483647) {
            for fi14 := 0; fi14 < 10; fi14++ {
                v10 = -238
                fmt.Println(v9, v9, fi14)
                v10 += (v10 >> ((255) & 63))
                _ = fi14
            }
            v15 := GenS1{F0: 793}
            v9 -= (^(v9))
            _ = v15
        } else {
            if (true && true) {
                v9 = (v9 >> ((180) & 63))
                v9 -= v9
            }
        }
    case "":
        v16 := []int{v9, (v9 - -92), (^(v9)), 9223372036854775807}
        _ = v16
    }
    _ = v10
    _ = v9
    return 42
}

func main() {
    var v1 uint16 = 255
    var v2 int8 = 11
    v3 := []string{"hello", (((("a" + "") + "") + "foo bar") + "hello"), "hello"}
    var v4 int16 = 1000
    v5 := make(map[int]uint32)
    var v6 int16 = v4
    v17 := GenS2{F0: 65535}
    fmt.Println(v1, v1)
    var v18 int = 32767
    for fi19 := 0; fi19 < 13; fi19++ {
        switch "foo bar" {
        case "":
            v4 -= (v4 % (((v4 | v4)) | 1))
        case "a":
            v2 += (v2 / ((v2) | 1))
        default:
            var v20 uint = genFn1()
            _ = v20
        }
        var v21 int8 = v2
        _ = fi19
        _ = v21
    }
    switch (v18 ^ v18) {
    case 4294967295:
        v22 := GenS1{F0: genFn1()}
        _ = v22
    case 9223372036854775807:
        v3 = append(v3, ("a" + ""))
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        for fi24 := 0; fi24 < 5; fi24++ {
            if false {
                v6 -= (v6 + v6)
                var v25 bool = true
                _ = v25
            } else {
                switch "hello" {
                case "a":
                    switch ("" + "a") {
                    case "\n\t":
                        switch (-(fi23)) {
                        case 2147483647:
                            v1 = (^(v1))
                            var v26 int64 = 365
                            _ = v26
                        case 749:
                            v2 += v2
                        }
                    case "a":
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                    default:
                        if (1 >= 9223372036854775807) {
                            var v27 int32 = -2147483648
                            _ = v27
                        } else {
                            v3 = append(v3, "foo bar")
                        }
                    }
                default:
                    switch 524 {
                    case -183:
                        var v28 int16 = 0
                        _ = v28
                    case -1:
                        var v29 uint32 = 0
                        _ = v29
                    case -650:
                        var v30 int8 = (v2 / ((v2) | 1))
                        _ = v30
                    default:
                        v31 := GenS2{F0: 65535}
                        v18 -= (fi24 >> ((2) & 63))
                        v1 = v1
                        _ = v31
                    }
                }
                v32 := GenS1{F0: 42}
                if (genFn1() != genFn1()) {
                    var v33 string = "\n\t"
                    v33 = "hello"
                    _ = v33
                }
                _ = v32
            }
            v34 := v5[fi24]
            _ = len(v3)
            _ = fi24
            _ = v34
        }
        _ = fi23
    }
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    var v35 int = (v18 | v18)
    v2 = (v2 * (v2 >> ((0) & 63)))
    var v36 int64 = -1
    fmt.Println(v4, v1, v1)
    v37 := GenS2{F0: 823}
    for fi38 := 0; fi38 < 13; fi38++ {
        fmt.Println(v6, v18, v36)
        _ = fi38
    }
    if len(v3) > 0 {
        v3[0] = "a"
    }
    v39 := v5[(v18 | v35)]
    fmt.Println(v6, v39, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
