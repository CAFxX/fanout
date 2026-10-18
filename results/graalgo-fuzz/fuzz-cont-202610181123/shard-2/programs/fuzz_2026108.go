package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int16 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int64; F1 uint32; F2 int16 }
func genFn0() int64 {
    for fi5 := 0; fi5 < 17; fi5++ {
        if true {
            v6 := []int{fi5, (fi5 / (((^(fi5))) | 1)), (fi5 / (((-(fi5))) | 1)), fi5}
            var v7 uint32 = 947
            _ = v6
            _ = v7
        }
        v8 := GenS1{F0: 1, F1: (-(fi5)), F2: 1}
        var v9 int8 = 127
        _ = fi5
        _ = v8
        _ = v9
    }
    var v10 float64 = float64((1.5 / 1.5))
    _ = v10
    return 9223372036854775807
}
func genFn1(p0 uint32, p1 int8, p2 uint64) int64 {
    fmt.Println(p0)
    for fi11 := 0; fi11 < 16; fi11++ {
        fmt.Println(p0)
        switch ("a" + "\n\t") {
        case "\n\t":
            if (("hello" <= "hello") && false) {
                fmt.Println(p2, p2)
                v12 := GenS1{F0: (^(p0)), F1: -1, F2: 0}
                fmt.Println(p2)
                _ = v12
            }
        case "":
            for fi13 := 0; fi13 < 17; fi13++ {
                if ((65 != (^(p1))) || (1 >= 255)) {
                    v14 := []int{fi11, fi13, (^(fi13)), fi13}
                    if ((p0 * 2147483647) > (p0 ^ 2147483647)) {
                        p1 += (p1 + 117)
                        p1 -= p1
                        var v15 int64 = 1
                        _ = v15
                    } else {
                        v16 := GenS1{F0: 42, F1: 0, F2: -32768}
                        p0 = p0
                        var v17 float32 = -2.25
                        _ = v16
                        _ = v17
                    }
                    var v18 uint = 0
                    _ = v14
                    _ = v18
                }
                p1 = p1
                fmt.Println(fi13)
                _ = fi13
            }
        case "foo bar":
            v19 := []int{(fi11 % ((0) | 1))}
            _ = v19
        }
        fmt.Println(p0, p0, p0)
        _ = fi11
    }
    fmt.Println(p2)
    v20 := GenS1{F0: (p0 * p0), F1: -1, F2: 776}
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return -764
}

func main() {
    v1 := []string{"a", "a", "hello"}
    var v2 int32 = 696
    var v3 string = "hello"
    v4 := make(map[int]string)
    var v21 int = -2147483648
    v21 = (v21 << ((308) & 63))
    if false {
        var v22 float64 = 3.14159
        fmt.Println(v2)
        _ = v22
    } else {
        var v23 uint = 817
        _ = v23
    }
    v3 = v3
    var v24 uint8 = 201
    var v25 float64 = 1e10
    var v26 int = (v21 >> ((315) & 63))
    var v27 int64 = genFn1(65535, 44, 1)
    for fi28 := 0; fi28 < 18; fi28++ {
        for fi29 := 0; fi29 < 16; fi29++ {
            switch ("foo bar" + "a") {
            case "hello":
                var v30 uint8 = v24
                _ = v30
            }
            _ = fi29
        }
        switch (v21 * 9223372036854775807) {
        case 42:
            switch v21 {
            case -9223372036854775808:
                v31 := GenS2{F0: (v24 - (v24 + v24)), F1: fi28}
                _ = v31
            }
        }
        if false {
            v32 := GenS3{F0: (v27 - (-(v27))), F1: 171, F2: 0}
            switch "\n\t" {
            case "a":
                v27 = (v27 % (((v27 - v27)) | 1))
            case "foo bar":
                v33 := v4[(fi28 - v21)]
                _ = v33
            case "\n\t":
                v4[(v26 >> ((65535) & 63))] = v3
            }
            var v34 uint = 18446744073709551615
            _ = v32
            _ = v34
        } else {
            var v35 int64 = genFn0()
            for fi36 := 0; fi36 < 1; fi36++ {
                v37 := GenS3{F0: (v35 >> ((697) & 63)), F1: 42, F2: 0}
                _ = fi36
                _ = v37
            }
            _ = v35
        }
        _ = fi28
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    delete(v4, v26)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
