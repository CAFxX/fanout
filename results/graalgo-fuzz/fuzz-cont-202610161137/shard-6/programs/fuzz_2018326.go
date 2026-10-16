package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 bool; F1 bool; F2 string }
type GenS5 struct { F0 string; F1 int8; F2 uint64 }

func main() {
    var v1 uint8 = 1
    v2 := []uint32{4294967295, 4294967295, 209}
    var v3 uint = uint(v1)
    var v4 uint64 = 42
    v5 := []string{(("" + "") + "a"), ("a" + ("\n\t" + "\n\t"))}
    v6 := make(map[string]int)
    v7 := GenS1{F0: "", F1: 1}
    if len(v5) > 0 {
        v5[0] = "hello"
    }
    var v8 int8 = -1
    var v9 string = ("\n\t" + "hello")
    var v10 uint64 = (v4 << ((v3) & 63))
    if (v8 != v8) {
        fmt.Println(v8, v1)
    }
    var v11 string = v9
    if (true && true) {
        for fi12 := 0; fi12 < 11; fi12++ {
            v13 := GenS2{F0: (v8 % ((int8(fi12)) | 1))}
            if (true && (!(0 < 255))) {
                for fi14 := 0; fi14 < 16; fi14++ {
                    var v15 int16 = -1
                    v10 -= (v4 + (v4 / ((v10) | 1)))
                    v3 = v3
                    _ = fi14
                    _ = v15
                }
                var v16 uint = (^(v3))
                _ = v16
            }
            if len(v5) > 0 {
                v5[0] = ("" + v11)
            }
            _ = fi12
            _ = v13
        }
        if (1 == (v8 >> ((0) & 63))) {
            if true {
                var v17 int8 = (^(v8))
                if ((v11 <= ((v11 + v11) + v9)) || (0 > 110)) {
                    for fi18 := 0; fi18 < 18; fi18++ {
                        switch (v11 + v11) {
                        case "hello":
                            _ = len(v6)
                        case "":
                            v19 := GenS3{F0: v11, F1: (v1 / ((128) | 1))}
                            _ = v19
                        }
                        _ = fi18
                    }
                } else {
                    v4 = v10
                    if false {
                        v20 := GenS2{F0: (v17 & v8)}
                        if false {
                            v21 := GenS3{F0: "\n\t", F1: v1}
                            _ = v21
                        } else {
                            if false {
                                switch 32767 {
                                case 947:
                                    var v22 uint64 = (^(v10))
                                    _ = v22
                                case 720:
                                    v23 := GenS3{F0: (v11 + ((v9 + (("hello" + v9) + v9)) + ("foo bar" + v11))), F1: v1}
                                    _ = v23
                                case 0:
                                    if (-2147483648 <= 0) {
                                        v24 := v6["a"]
                                        _ = v24
                                    }
                                }
                                for fi25 := 0; fi25 < 2; fi25++ {
                                    _ = len(v5)
                                    v5 = append(v5, v11)
                                    _ = fi25
                                }
                                var v26 int = 127
                                _ = v26
                            }
                            var v27 uint32 = 42
                            _ = v27
                        }
                        _ = v20
                    } else {
                        v3 = v3
                        for fi28 := 0; fi28 < 19; fi28++ {
                            v8 -= (v17 ^ v8)
                            var v29 int64 = -927
                            fmt.Println(v1)
                            _ = fi28
                            _ = v29
                        }
                        v30 := GenS4{F0: (4294967295 <= 4294967295), F1: true, F2: "\n\t"}
                        _ = v30
                    }
                    v10 = (v10 * v4)
                }
                delete(v6, "hello")
                _ = v17
            } else {
                var v31 uint = v3
                _ = v31
            }
            v6[v9] = -668
        } else {
            for fi32 := 0; fi32 < 4; fi32++ {
                v4 -= (v4 | v4)
                _ = fi32
            }
            v33 := GenS1{F0: "", F1: 609}
            _ = v33
        }
    }
    var v34 string = (v9 + v9)
    var v35 float32 = -2.25
    fmt.Println(v10, v10)
    for fi36 := 0; fi36 < 3; fi36++ {
        for fi37 := 0; fi37 < 11; fi37++ {
            var v38 float64 = 0.0
            if false {
                _ = len(v6)
                var v39 float32 = (3.14159 + v35)
                _ = v39
            } else {
                v40 := GenS2{F0: (v8 & (^(v8)))}
                v41 := GenS5{F0: v34, F1: v8, F2: v4}
                _ = v40
                _ = v41
            }
            _ = fi37
            _ = v38
        }
        var v42 uint16 = 32767
        _ = fi36
        _ = v42
    }
    var v43 uint = (v3 + (v3 / ((v3) | 1)))
    var v44 uint64 = v10
    var v45 int16 = 716
    var v46 string = v34
    if len(v5) > 0 {
        v5[0] = "\n\t"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
