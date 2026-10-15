package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0() int64 {
    var v6 bool = (true && true)
    v6 = (32767 > 32767)
    var v7 bool = true
    v8 := []int{42, -9223372036854775808, 42}
    _ = v6
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 int64) uint64 {
    var v9 int16 = -43
    var v10 bool = (((p0 >> ((465) & 63)) != (-(p0))) && ((true || (false || false)) || ((4 == 0) || (!(1.5 >= 3.14159)))))
    switch "foo bar" {
    case "a":
        v9 = -142
    case "foo bar":
        var v11 float64 = ((0.1 * 3.14159) - float64(((float32((((0.0 - 1.5) - 0.1) + 0.0)) + (1.5 / 0.1)) / ((3.14159 * 3.14159) / 1e10))))
        _ = v11
    default:
        v12 := []int{0}
        _ = v12
    }
    v9 -= v9
    _ = p0
    _ = v10
    _ = v9
    return 1
}

func main() {
    var v1 uint8 = 2
    var v2 string = "hello"
    var v3 float64 = 3.14159
    v4 := make(map[string]int)
    v5 := []uint32{0, 0, 2147483647}
    if ((((9223372036854775807 == genFn0()) || (72 < v1)) || true) && false) {
        for fi13 := 0; fi13 < 4; fi13++ {
            for fi14 := 0; fi14 < 1; fi14++ {
                var v15 float32 = 1e10
                switch fi13 {
                case 32767:
                    v16 := GenS1{F0: ""}
                    _ = v16
                default:
                    v4["a"] = fi14
                }
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
        v2 = "hello"
    } else {
        for fi17 := 0; fi17 < 2; fi17++ {
            var v18 int32 = -993
            v19 := GenS1{F0: "a"}
            switch 177 {
            case 32768:
                fmt.Println(v2, v1, v18)
            case 42:
                fmt.Println(v3, v18)
            }
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    if (756 == 42) {
        v3 = 1.5
        var v20 uint64 = 4294967295
        _ = v20
    }
    fmt.Println(v2)
    v2 = v2
    var v21 int32 = -1
    if (!true) {
        delete(v4, (("" + v2) + (((v2 + "hello") + "\n\t") + v2)))
        v2 = ""
    } else {
        if (!true) {
            switch v2 {
            case "a":
                if true {
                    var v22 uint64 = 1
                    var v23 float32 = 1e10
                    var v24 int = 171
                    _ = v22
                    _ = v23
                    _ = v24
                } else {
                    var v25 int = 42
                    for fi26 := 0; fi26 < 13; fi26++ {
                        fmt.Println(v25, v2)
                        if false {
                            var v27 bool = (((119 < -63) && false) || true)
                            if (0.0 <= (0.1 + 1e10)) {
                                v28 := GenS1{F0: (v2 + (((("hello" + "foo bar") + ("" + v2)) + ((("\n\t" + v2) + (v2 + (v2 + v2))) + ((("hello" + "\n\t") + v2) + "foo bar"))) + (v2 + v2)))}
                                v27 = v27
                                switch v2 {
                                case "hello":
                                    v27 = ((v27 && false) || v27)
                                case "foo bar":
                                    v25 = v25
                                }
                                _ = v28
                            } else {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                var v29 int = fi26
                                _ = v29
                            }
                            _ = v27
                        }
                        if false {
                            switch (v2 + (v2 + v2)) {
                            case "":
                                var v30 uint = 65535
                                _ = v30
                            case "foo bar":
                                var v31 bool = true
                                _ = v31
                            case "hello":
                                for fi32 := 0; fi32 < 1; fi32++ {
                                    v1 -= 6
                                    switch (v2 + "foo bar") {
                                    case "hello":
                                        var v33 uint8 = (v1 >> ((402) & 63))
                                        _ = v33
                                    case "\n\t":
                                        _ = len(v4)
                                    }
                                    var v34 int32 = v21
                                    _ = fi32
                                    _ = v34
                                }
                            }
                            fmt.Println(v1)
                        }
                        _ = fi26
                    }
                    _ = v25
                }
            case "\n\t":
                v1 = (v1 ^ v1)
            }
            if false {
                var v35 float32 = (3.14159 / 1.5)
                v36 := GenS1{F0: (v2 + (("a" + v2) + "hello"))}
                _ = v35
                _ = v36
            } else {
                var v37 int = 0
                _ = v37
            }
        } else {
            for fi38 := 0; fi38 < 20; fi38++ {
                v39 := GenS1{F0: v2}
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi38
                _ = v39
            }
            var v40 string = v2
            v41 := GenS1{F0: v2}
            _ = v40
            _ = v41
        }
    }
    delete(v4, "foo bar")
    switch (("foo bar" + "hello") + v2) {
    case "\n\t":
        v42 := GenS1{F0: "\n\t"}
        _ = v42
    }
    fmt.Println(v1, v1, v1)
    v1 = v1
    var v43 int16 = 1000
    var v44 uint = 0
    var v45 uint = v44
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}
