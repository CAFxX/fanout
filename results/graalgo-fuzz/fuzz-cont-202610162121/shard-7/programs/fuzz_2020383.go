package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 uint8; F2 int }
type GenS4 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS5 struct { F0 int; F1 uint8 }
func genFn0(p0 uint64, p1 uint) int16 {
    for fi5 := 0; fi5 < 1; fi5++ {
        if true {
            if (2147483647 < -2147483648) {
                var v6 int64 = 453
                var v7 uint8 = 127
                v8 := GenS1{F0: 0}
                _ = v6
                _ = v7
                _ = v8
            }
            v9 := []int{}
            var v10 uint64 = (p0 - (p0 | 42))
            _ = v10
            _ = v9
        }
        v11 := GenS1{F0: 255}
        if true {
            for fi12 := 0; fi12 < 20; fi12++ {
                p1 = p1
                switch "hello" {
                case "\n\t":
                    switch fi12 {
                    case -2147483648:
                        v13 := []int{fi5}
                        _ = v13
                    case 32768:
                        if ((p0 + p0) == p0) {
                            v14 := []int{(^(fi5)), (fi12 << ((p1) & 63)), (fi5 / (((^(fi5))) | 1))}
                            p1 -= (^(p1))
                            fmt.Println(p1, fi5, fi5)
                            _ = v14
                        }
                    case -9223372036854775808:
                        p0 = (p0 - (^(p0)))
                        v15 := GenS2{F0: 32767}
                        var v16 int64 = 540
                        _ = v15
                        _ = v16
                    }
                }
                _ = fi12
            }
        } else {
            fmt.Println(fi5)
        }
        _ = fi5
        _ = v11
    }
    if true {
        p1 = 1
        p1 -= p1
    }
    v17 := []int{4294967295, 128}
    fmt.Println(p1, p1)
    v18 := GenS2{F0: -1}
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 42
}

func main() {
    v1 := []int{-320}
    v2 := make(map[int]bool)
    var v3 int32 = -248
    v4 := []int{-274}
    var v19 int64 = -474
    v20 := GenS1{F0: 65535}
    v21 := v2[607]
    var v22 int16 = int16(v3)
    v3 -= v3
    if v21 {
        var v23 string = (("\n\t" + ("hello" + "a")) + "\n\t")
        for fi24 := 0; fi24 < 18; fi24++ {
            v25 := v2[(fi24 ^ fi24)]
            _ = fi24
            _ = v25
        }
        _ = v23
    } else {
        if true {
            v26 := GenS3{F0: 0, F1: 128, F2: 9223372036854775807}
            _ = v26
        } else {
            switch int(v3) {
            case -241:
                v27 := GenS4{F0: 865, F1: v19, F2: 941}
                _ = v27
            }
            fmt.Println(v21, v21, v21)
        }
        for fi28 := 0; fi28 < 11; fi28++ {
            v29 := GenS1{F0: 42}
            v30 := GenS5{F0: fi28, F1: 42}
            for fi31 := 0; fi31 < 4; fi31++ {
                var v32 uint = 0
                _ = fi31
                _ = v32
            }
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    v22 += genFn0(43, 833)
    v33 := GenS2{F0: genFn0(1, 334)}
    switch -1 {
    case 513:
        if (!(!(v21 && v21))) {
            var v34 int8 = 100
            _ = v34
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    case 32768:
        for fi35 := 0; fi35 < 14; fi35++ {
            v22 -= 32767
            v22 -= v22
            _ = fi35
        }
    case -185:
        switch ("" + "hello") {
        case "\n\t":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "a":
            fmt.Println(v3, v22)
        case "":
            switch ("hello" + "") {
            case "\n\t":
                v3 = (^(v3))
            }
        }
    }
    fmt.Println(v22)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v33
    _ = v4
}
