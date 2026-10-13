package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 uint32; F2 int8 }
type GenS3 struct { F0 int }
func genFn0() bool {
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := []int{}
        var v9 int = fi7
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 4294967295}
    if false {
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 uint64 = 221
            switch ("hello" + "\n\t") {
            case "foo bar":
                v12 = 42
            case "a":
                v12 = (v12 >> ((42) & 63))
            case "\n\t":
                if ((1e10 * (0.1 + (0.0 + 3.14159))) != (1e10 / (((0.1 + 1.5) + 1.5) / (1e10 + 1.5)))) {
                    var v13 bool = (9223372036854775807 <= 9223372036854775807)
                    if false {
                        if (!(-1 <= 0)) {
                            var v14 int32 = 1
                            var v15 int32 = (v14 << ((uint(v12)) & 63))
                            _ = v14
                            _ = v15
                        }
                        v16 := GenS1{F0: v12}
                        _ = v16
                    }
                    v13 = true
                    _ = v13
                } else {
                    var v17 string = (("" + "\n\t") + "foo bar")
                    for fi18 := 0; fi18 < 18; fi18++ {
                        v19 := GenS2{F0: 42, F1: 2147483647, F2: 127}
                        var v20 string = (("foo bar" + "foo bar") + v17)
                        v12 -= v12
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 10; fi21++ {
                        switch (((v17 + "a") + v17) + (v17 + v17)) {
                        case "\n\t":
                            switch (fi11 | fi11) {
                            case 42:
                                fmt.Println(fi11, fi21)
                            }
                        case "a":
                            v22 := []int{(-(fi21))}
                            _ = v22
                        }
                        v12 -= v12
                        _ = fi21
                    }
                    _ = v17
                }
            default:
                for fi23 := 0; fi23 < 16; fi23++ {
                    v24 := []int{42}
                    v25 := GenS1{F0: v12}
                    fmt.Println(v12)
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
            _ = fi11
            _ = v12
        }
        var v26 uint16 = 0
        var v27 uint32 = 610
        _ = v26
        _ = v27
    }
    _ = v10
    return true
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]int)
    var v3 float32 = ((0.0 * 3.14159) / 1.0)
    v4 := make(map[int]bool)
    var v5 bool = true
    var v6 uint64 = 0
    fmt.Println(v5, v3, v6)
    if true {
        delete(v1, 65535)
        for fi28 := 0; fi28 < 14; fi28++ {
            v5 = (191 > 1)
            v29 := GenS1{F0: v6}
            v30 := []int{}
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    v5 = (255 <= 255)
    var v31 uint64 = (v6 ^ 1)
    for fi32 := 0; fi32 < 3; fi32++ {
        v33 := []int{}
        v34 := GenS3{F0: fi32}
        if (!false) {
            var v35 uint8 = 127
            fmt.Println(v35, v3)
            for fi36 := 0; fi36 < 3; fi36++ {
                var v37 bool = v5
                var v38 uint32 = 762
                _ = fi36
                _ = v37
                _ = v38
            }
            _ = v35
        } else {
            if v5 {
                v39 := GenS3{F0: (fi32 % ((-316) | 1))}
                _ = v39
            } else {
                v31 += (v6 - (v6 & (v6 << ((18446744073709551615) & 63))))
                switch ((("hello" + (("a" + ("\n\t" + (((("foo bar" + ("a" + "\n\t")) + ("\n\t" + "\n\t")) + ("\n\t" + "hello")) + "foo bar"))) + "\n\t")) + "a") + ("a" + "hello")) {
                case "hello":
                    fmt.Println(v3, fi32, v6)
                }
            }
            for fi40 := 0; fi40 < 9; fi40++ {
                var v41 float32 = v3
                _ = fi40
                _ = v41
            }
        }
        _ = fi32
        _ = v33
        _ = v34
    }
    var v42 uint16 = 1
    v6 -= (v6 + v6)
    v3 -= float32(1.5)
    for fi43 := 0; fi43 < 12; fi43++ {
        var v44 int8 = 30
        var v45 int16 = -1
        _ = fi43
        _ = v44
        _ = v45
    }
    v3 = (1e10 + -2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
