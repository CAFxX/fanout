package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int }

func main() {
    var v1 int16 = 1
    var v2 uint8 = 177
    var v3 bool = true
    v4 := []int{1, 32767}
    v5 := make(map[string]bool)
    v6 := GenS1{F0: ("a" + "hello"), F1: 20, F2: "hello"}
    fmt.Println(v2)
    v7 := GenS2{F0: ("foo bar" + ("\n\t" + "foo bar"))}
    switch 32767 {
    case 0:
        switch (("\n\t" + "hello") + "a") {
        case "":
            v8 := v5["hello"]
            _ = v8
        case "\n\t":
            v9 := GenS3{F0: 591}
            _ = v9
        }
    default:
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 float32 = 3.14159
            var v12 float64 = ((-2.25 / 1e10) * -2.25)
            v13 := GenS1{F0: ("hello" + "foo bar"), F1: 0, F2: ("\n\t" + "a")}
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        switch fi14 {
        case 352:
            if v3 {
                for fi15 := 0; fi15 < 10; fi15++ {
                    var v16 bool = (v3 && false)
                    if (-131 == -9223372036854775808) {
                        var v17 uint16 = 2
                        _ = v17
                    } else {
                        switch -836 {
                        case 920:
                            fmt.Println(v16, v3)
                        case 32767:
                            var v18 float32 = 0.1
                            _ = v18
                        case 127:
                            delete(v5, (("a" + "foo bar") + "hello"))
                        }
                        v3 = (true && (3.14159 <= float64(0.1)))
                        var v19 uint8 = v2
                        fmt.Println(v2, v1, v1)
                        _ = v19
                    }
                    _ = fi15
                    _ = v16
                }
                for fi20 := 0; fi20 < 5; fi20++ {
                    v21 := GenS4{F0: 1}
                    _ = fi20
                    _ = v21
                }
            }
        }
        if (int8(v1) < 100) {
            var v22 uint = 18446744073709551615
            _ = v22
        }
        _ = fi14
    }
    if true {
        var v23 int = 0
        v1 += (-(v1))
        _ = v23
    } else {
        v2 -= v2
        var v24 uint16 = 202
        if (-453 != 42) {
            delete(v5, "")
            v1 = v1
            if true {
                for fi25 := 0; fi25 < 1; fi25++ {
                    v5["a"] = v3
                    v2 = v2
                    _ = fi25
                }
                var v26 float32 = (0.0 * 1e10)
                fmt.Println(v3, v1)
                _ = v26
            } else {
                for fi27 := 0; fi27 < 4; fi27++ {
                    v1 = (^(v1))
                    v28 := GenS5{F0: 9223372036854775807}
                    for fi29 := 0; fi29 < 18; fi29++ {
                        v1 = int16(v2)
                        if false {
                            if v3 {
                                v24 = (v24 * v24)
                                var v30 int = fi29
                                fmt.Println(v3, v1, fi29)
                                if (v30 > fi29) {
                                    var v31 uint8 = v2
                                    delete(v5, "foo bar")
                                    _ = v31
                                } else {
                                    fmt.Println(v30, v3)
                                    delete(v5, "a")
                                    var v32 uint32 = 65535
                                    _ = v32
                                }
                                fmt.Println(v2, fi27, v2)
                                _ = v30
                            } else {
                                v33 := GenS6{F0: (fi29 + (fi27 | (^(fi27))))}
                                _ = len(v5)
                                fmt.Println(v2)
                                _ = v33
                            }
                        }
                        v34 := v5[("" + ("a" + "foo bar"))]
                        v1 = (v1 + 337)
                        _ = fi29
                        _ = v34
                    }
                    _ = fi27
                    _ = v28
                }
                var v35 uint64 = 42
                delete(v5, ((("" + "") + "hello") + ("hello" + "foo bar")))
                _ = v35
            }
        }
        _ = v24
    }
    fmt.Println(v1, v2, v3)
    var v36 int64 = 77
    fmt.Println(v2, v2)
    v37 := v5[("hello" + ("a" + "foo bar"))]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
