package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    var v3 uint64 = 1
    var v4 bool = true
    var v5 int64 = -733
    var v6 int16 = 760
    if true {
        switch -17 {
        case 2147483647:
            for fi7 := 0; fi7 < 16; fi7++ {
                v5 -= v5
                if (v4 || v4) {
                    v4 = (42 != v3)
                }
                _ = fi7
            }
        case 127:
            delete(v1, 42)
        }
        var v8 uint8 = 220
        _ = v8
    } else {
        v9 := GenS1{F0: 651}
        var v10 int64 = v5
        var v11 int64 = v10
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 bool = v4
    v13 := GenS1{F0: 2}
    v5 = 0
    v2[32768] = 4294967295
    if (v6 != (v6 | 465)) {
        if v4 {
            v4 = v4
            var v14 float64 = 1.5
            _ = v14
        }
        switch (("a" + "hello") + "hello") {
        case "\n\t":
            var v15 string = "hello"
            _ = v15
        }
        if ((1e10 / 1.5) < float64((1e10 - 0.1))) {
            var v16 float32 = ((1e10 / -2.25) - 1.5)
            v17 := GenS2{F0: (v5 * -9223372036854775808)}
            _ = v16
            _ = v17
        }
    } else {
        var v18 int64 = -9223372036854775808
        if (!v4) {
            fmt.Println(v12, v18)
        }
        _ = v18
    }
    v19 := v1[834]
    var v20 uint64 = (^(v3))
    if v19 {
        v4 = (128 != 1)
    } else {
        var v21 int = 0
        _ = v21
    }
    if true {
        v4 = v4
    }
    if v4 {
        switch "a" {
        case "":
            for fi22 := 0; fi22 < 13; fi22++ {
                if true {
                    v23 := GenS2{F0: (v5 & v5)}
                    fmt.Println(v3)
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 9; fi24++ {
                        fmt.Println(v6)
                        _ = fi24
                    }
                    if v4 {
                        v19 = (!false)
                    } else {
                        switch "a" {
                        case "foo bar":
                            v5 = (v5 | 42)
                        case "\n\t":
                            v20 += (v20 + v20)
                        case "hello":
                            for fi25 := 0; fi25 < 18; fi25++ {
                                if true {
                                    var v26 float64 = 0.1
                                    v26 += (v26 + v26)
                                    _ = v26
                                } else {
                                    v27 := []int{(fi22 - fi22)}
                                    var v28 int32 = -1
                                    if false {
                                        switch "foo bar" {
                                        case "a":
                                            if false {
                                                v29 := v1[850]
                                                var v30 int32 = 1
                                                var v31 float32 = (-2.25 - ((((-2.25 + float32(3.14159)) * -2.25) + -2.25) / (3.14159 / 0.1)))
                                                _ = v29
                                                _ = v30
                                                _ = v31
                                            }
                                        }
                                        v32 := GenS2{F0: (v5 + v5)}
                                        _ = v32
                                    } else {
                                        v12 = v4
                                    }
                                    if v4 {
                                        var v33 uint = 1
                                        _ = v33
                                    } else {
                                        v20 = (v20 | v20)
                                        v6 = v6
                                        var v34 uint32 = 4294967295
                                        _ = v34
                                    }
                                    _ = v27
                                    _ = v28
                                }
                                var v35 float64 = (0.1 / 0.1)
                                _ = fi25
                                _ = v35
                            }
                        }
                    }
                    if v19 {
                        var v36 float32 = (-2.25 * 0.0)
                        v37 := GenS1{F0: 488}
                        v38 := GenS1{F0: 65535}
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                }
                var v39 uint8 = 2
                v40 := GenS1{F0: 1}
                _ = fi22
                _ = v39
                _ = v40
            }
        case "foo bar":
            v41 := GenS3{F0: v3}
            _ = v41
        case "hello":
            v42 := v2[2147483647]
            _ = v42
        default:
            fmt.Println(v6)
        }
    }
    var v43 int32 = -207
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
