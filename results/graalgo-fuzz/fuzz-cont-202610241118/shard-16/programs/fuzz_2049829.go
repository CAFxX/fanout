package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 uint64) int16 {
    fmt.Println(p0)
    var v4 int32 = 467
    if (!true) {
        v5 := GenS1{F0: uint16(p0)}
        v6 := []int{42, 127, -590, 42}
        _ = v5
        _ = v6
    }
    _ = p0
    _ = v4
    return -497
}
func genFn1(p0 uint64, p1 int) string {
    v7 := GenS1{F0: 1}
    v8 := GenS1{F0: 32767}
    if false {
        v9 := []int{(p1 / ((-986) | 1))}
        _ = v9
    }
    switch "hello" {
    case "hello":
        if (!true) {
            for fi10 := 0; fi10 < 13; fi10++ {
                v11 := []int{17, fi10}
                var v12 int = (fi10 ^ fi10)
                _ = fi10
                _ = v11
                _ = v12
            }
        } else {
            p1 = p1
            for fi13 := 0; fi13 < 17; fi13++ {
                v14 := []int{(-(fi13))}
                if (811 > 1) {
                    p1 = p1
                }
                switch (^(p1)) {
                case 946:
                    for fi15 := 0; fi15 < 1; fi15++ {
                        fmt.Println(fi15, fi13, p1)
                        _ = fi15
                    }
                case 127:
                    var v16 float32 = (-2.25 + -2.25)
                    _ = v16
                case 875:
                    var v17 uint = 173
                    _ = v17
                }
                _ = fi13
                _ = v14
            }
            for fi18 := 0; fi18 < 1; fi18++ {
                if false {
                    var v19 string = "a"
                    for fi20 := 0; fi20 < 4; fi20++ {
                        v21 := []int{-814, (fi18 / ((fi20) | 1))}
                        v22 := GenS1{F0: 42}
                        if true {
                            var v23 uint16 = 255
                            var v24 int = (fi18 % ((p1) | 1))
                            _ = v23
                            _ = v24
                        } else {
                            fmt.Println(v19)
                        }
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    v25 := GenS1{F0: 2}
                    _ = v19
                    _ = v25
                }
                p0 = (p0 >> ((866) & 63))
                var v26 int16 = 1000
                _ = fi18
                _ = v26
            }
        }
    case "foo bar":
        var v27 float64 = (0.1 - -2.25)
        _ = v27
    case "\n\t":
        var v28 int8 = 1
        _ = v28
    }
    for fi29 := 0; fi29 < 15; fi29++ {
        var v30 float32 = 1e10
        p0 = p0
        _ = fi29
        _ = v30
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return ""
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[string]int)
    v3 := make(map[string]uint32)
    if (446 >= 0) {
        v31 := []int{}
        _ = v31
    }
    v1 += (v1 << ((0) & 63))
    for fi32 := 0; fi32 < 16; fi32++ {
        v33 := []int{(-(fi32)), 1, (fi32 >> ((402) & 63))}
        _ = fi32
        _ = v33
    }
    switch ("hello" + "hello") {
    case "":
        var v34 uint16 = 1
        _ = v34
    case "\n\t":
        v1 = (v1 - (v1 ^ (v1 + v1)))
    }
    fmt.Println(v1)
    fmt.Println(v1)
    switch 42 {
    case 32767:
        _ = len(v3)
    }
    v1 = 0
    v35 := []int{1, 1, 0, 4294967295}
    for fi36 := 0; fi36 < 5; fi36++ {
        if len(v35) >= 2 {
            v35 = v35[1:len(v35)]
        }
        _ = fi36
    }
    fmt.Println(v1)
    if len(v35) >= 2 {
        v35 = v35[1:len(v35)]
    }
    var v37 uint8 = 2
    fmt.Println(v1, v37)
    var v38 bool = false
    v37 = (v37 ^ 127)
    var v39 bool = (!false)
    fmt.Println(v37)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v38
    _ = v39
}
