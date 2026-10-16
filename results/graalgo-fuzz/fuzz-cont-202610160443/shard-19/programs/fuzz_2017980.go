package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint64 }
type GenS2 struct { F0 int }
func genFn0(p0 bool, p1 uint16, p2 bool) uint64 {
    if p0 {
        var v4 float32 = -2.25
        v5 := []int{}
        _ = v4
        _ = v5
    } else {
        fmt.Println(p2, p0, p0)
        var v6 uint32 = 2
        fmt.Println(v6, p0, v6)
        _ = v6
    }
    var v7 string = "foo bar"
    var v8 int8 = 0
    var v9 int16 = -32768
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[uint32]int)
    v3 := make(map[int]uint32)
    if (42 <= -303) {
        if (574 > 4294967295) {
            v10 := v3[-979]
            v1 = v1
            _ = v10
        }
        v11 := []int{2147483647, 627, 127}
        _ = v11
    }
    if true {
        if (0 < 188) {
            var v12 uint64 = (v1 << ((255) & 63))
            switch ("\n\t" + "hello") {
            case "hello":
                fmt.Println(v1, v12)
            case "a":
                v13 := []int{-1, 1, 662}
                _ = v13
            }
            v12 -= (v12 % ((v12) | 1))
            _ = v12
        } else {
            v1 = 951
            var v14 float64 = 3.14159
            v15 := GenS1{F0: -1, F1: true, F2: v1}
            _ = v14
            _ = v15
        }
        var v16 int32 = 84
        delete(v2, 770)
        _ = v16
    } else {
        v17 := GenS1{F0: -9223372036854775808, F1: false, F2: v1}
        v18 := GenS2{F0: 206}
        for fi19 := 0; fi19 < 18; fi19++ {
            v20 := GenS2{F0: fi19}
            var v21 uint64 = genFn0(true, 32767, (!(42 < -213)))
            for fi22 := 0; fi22 < 5; fi22++ {
                if (true && (42 <= (^(v21)))) {
                    var v23 int8 = -1
                    _ = v23
                }
                var v24 int16 = 32767
                _ = fi22
                _ = v24
            }
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = v17
        _ = v18
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        for fi26 := 0; fi26 < 1; fi26++ {
            var v27 float64 = 3.14159
            _ = fi26
            _ = v27
        }
        for fi28 := 0; fi28 < 9; fi28++ {
            switch (fi25 << ((4294967295) & 63)) {
            case 32767:
                if true {
                    var v29 uint32 = 736
                    v29 = (v29 | (v29 | (v29 >> ((65535) & 63))))
                    v30 := GenS2{F0: (fi28 | 2147483647)}
                    _ = v29
                    _ = v30
                } else {
                    v31 := GenS2{F0: 362}
                    _ = v31
                }
            case 1:
                if true {
                    v32 := []int{(fi28 % (((-(fi25))) | 1))}
                    if false {
                        v33 := v3[(fi28 * fi28)]
                        fmt.Println(v1, fi25, fi28)
                        _ = v33
                    } else {
                        fmt.Println(fi28, fi28)
                        if (545 <= 1) {
                            v1 = (v1 - v1)
                            if ((("hello" + "a") > "a") || true) {
                                switch fi28 {
                                case 9223372036854775807:
                                    delete(v3, (fi25 + fi25))
                                case 32767:
                                    v34 := GenS2{F0: 4294967295}
                                    _ = v34
                                case 128:
                                    v35 := GenS2{F0: fi28}
                                    _ = v35
                                }
                            }
                            if false {
                                var v36 int8 = 127
                                _ = v36
                            } else {
                                var v37 string = "foo bar"
                                var v38 int8 = 98
                                var v39 uint16 = 0
                                _ = v37
                                _ = v38
                                _ = v39
                            }
                        }
                        var v40 int8 = -69
                        _ = v40
                    }
                    _ = v32
                } else {
                    var v41 int8 = 0
                    _ = v41
                }
            case -9223372036854775808:
                v42 := []int{(fi25 + (fi28 << ((730) & 63))), fi25, fi25}
                _ = v42
            default:
                v43 := GenS2{F0: fi25}
                _ = v43
            }
            _ = len(v2)
            _ = fi28
        }
        _ = fi25
    }
    v1 += genFn0(("a" < "foo bar"), 42, ((true && false) && false))
    v1 = genFn0(false, 185, false)
    v2[376] = 319
    fmt.Println(v1)
    var v44 uint32 = 2
    fmt.Println(v44, v1)
    var v45 string = "\n\t"
    if (0 <= (v44 % ((v44) | 1))) {
        fmt.Println(v45, v1, v1)
    } else {
        var v46 int32 = 1
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v45
}
