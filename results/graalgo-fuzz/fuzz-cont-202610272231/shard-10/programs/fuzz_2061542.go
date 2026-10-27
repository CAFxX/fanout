package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 uint32; F2 int }
func genFn0(p0 bool, p1 uint64) uint64 {
    p1 = (^(p1))
    p1 = p1
    fmt.Println(p1, p1)
    p0 = p0
    v4 := []int{32767, 0}
    _ = p0
    _ = p1
    _ = v4
    return (p1 + (p1 % (((p1 - p1)) | 1)))
}
func genFn1(p0 uint64, p1 int32) bool {
    var v5 int8 = int8(p0)
    v6 := GenS1{F0: 654, F1: -1, F2: 65535}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (-264 > 127)
}

func main() {
    var v1 int32 = -159
    var v2 float32 = 0.0
    v3 := []string{("a" + "hello"), (("hello" + "hello") + "a"), ("a" + ("hello" + "a"))}
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    var v7 int8 = -108
    var v8 uint64 = 42
    v1 = v1
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    var v9 bool = false
    v8 -= v8
    if (751 != 42) {
        var v10 uint32 = 2147483647
        _ = v10
    } else {
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 int32 = (^(v1))
            if (-9223372036854775808 == 9223372036854775807) {
                if v9 {
                    fmt.Println(v2, v12, v7)
                    v13 := GenS2{F0: 1}
                    var v14 string = "hello"
                    _ = v13
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 10; fi15++ {
                        _ = len(v3)
                        _ = fi15
                    }
                    var v16 int8 = v7
                    _ = v16
                }
            } else {
                var v17 bool = v9
                switch ("\n\t" + "foo bar") {
                case "a":
                    fmt.Println(v7, v7, v12)
                case "foo bar":
                    var v18 uint = uint(v8)
                    _ = v18
                case "\n\t":
                    switch (fi11 ^ 127) {
                    case -611:
                        v1 -= v1
                    case 32767:
                        v2 = (-2.25 * (v2 / v2))
                    case 42:
                        v17 = ((!(v8 != 1)) && false)
                    }
                }
                _ = v17
            }
            _ = fi11
            _ = v12
        }
    }
    var v19 uint32 = uint32(v8)
    v8 -= (^(v8))
    if v9 {
        v20 := GenS3{F0: 42, F1: (v19 & 4294967295), F2: -492}
        v2 += -2.25
        var v21 uint16 = 331
        _ = v20
        _ = v21
    } else {
        if v9 {
            var v22 uint8 = uint8(v8)
            var v23 string = "hello"
            if true {
                fmt.Println(v2, v2, v2)
            }
            _ = v22
            _ = v23
        }
        var v24 uint8 = 2
        for fi25 := 0; fi25 < 1; fi25++ {
            if v9 {
                _ = len(v3)
            } else {
                var v26 uint8 = 0
                if (v1 == (v1 ^ -2147483648)) {
                    for fi27 := 0; fi27 < 10; fi27++ {
                        if v9 {
                            v28 := GenS1{F0: 65535, F1: 0, F2: 42}
                            v1 -= v1
                            if true {
                                if true {
                                    var v29 int64 = -1
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    if v9 {
                                        fmt.Println(v1, v9)
                                    } else {
                                        var v30 string = "hello"
                                        _ = v30
                                    }
                                    _ = v29
                                } else {
                                    v31 := GenS1{F0: 304, F1: -32768, F2: 0}
                                    _ = v31
                                }
                                var v32 uint8 = (^(v26))
                                v26 = (v26 ^ v26)
                                _ = v32
                            }
                            fmt.Println(v26)
                            _ = v28
                        } else {
                            v33 := GenS2{F0: (^(v19))}
                            _ = v33
                        }
                        fmt.Println(v19)
                        if (758 >= 32767) {
                            if true {
                                v34 := GenS3{F0: 42, F1: (v19 ^ 2147483647), F2: (fi25 + int(v7))}
                                var v35 int = fi25
                                _ = v34
                                _ = v35
                            }
                        } else {
                            var v36 float32 = (((v2 / v2) + v2) + v2)
                            _ = v36
                        }
                        _ = fi27
                    }
                    var v37 int32 = (v1 - (^(v1)))
                    _ = v37
                } else {
                    if (210 < v19) {
                        var v38 int = (fi25 * fi25)
                        var v39 int32 = v1
                        _ = v38
                        _ = v39
                    }
                }
                _ = v26
            }
            v40 := GenS2{F0: (v19 ^ v19)}
            fmt.Println(fi25, fi25, v19)
            _ = fi25
            _ = v40
        }
        _ = v24
    }
    v1 += -782
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
