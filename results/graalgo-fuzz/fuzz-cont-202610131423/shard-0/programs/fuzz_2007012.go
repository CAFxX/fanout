package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 uint8 }
func genFn0(p0 int) uint8 {
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := []int{(-(fi6)), -1, p0}
        var v8 int = fi6
        _ = fi6
        _ = v7
        _ = v8
    }
    if false {
        switch ((("a" + ("\n\t" + "a")) + ("" + "")) + "\n\t") {
        case "hello":
            p0 += 0
        }
        if true {
            fmt.Println(p0)
            v9 := []int{p0, (p0 - p0)}
            _ = v9
        } else {
            p0 = p0
            for fi10 := 0; fi10 < 11; fi10++ {
                if (-2.25 < (-2.25 / float32(-2.25))) {
                    if (0.1 <= 0.1) {
                        var v11 int = fi10
                        v11 += 128
                        for fi12 := 0; fi12 < 15; fi12++ {
                            var v13 float64 = 0.0
                            _ = fi12
                            _ = v13
                        }
                        _ = v11
                    } else {
                        if true {
                            if false {
                                v14 := GenS1{F0: 65535}
                                var v15 int8 = 127
                                for fi16 := 0; fi16 < 7; fi16++ {
                                    v15 += v15
                                    _ = fi16
                                }
                                _ = v14
                                _ = v15
                            } else {
                                p0 += (fi10 ^ 127)
                                var v17 uint8 = 1
                                _ = v17
                            }
                        } else {
                            var v18 uint32 = 4294967295
                            v19 := []int{(fi10 << ((42) & 63)), (fi10 | (p0 / ((fi10) | 1))), (p0 & (fi10 >> ((0) & 63)))}
                            _ = v18
                            _ = v19
                        }
                    }
                } else {
                    var v20 string = "hello"
                    for fi21 := 0; fi21 < 5; fi21++ {
                        if (904 < 959) {
                            p0 = fi10
                            v22 := GenS1{F0: 42}
                            _ = v22
                        }
                        if false {
                            var v23 uint8 = 133
                            _ = v23
                        }
                        if (!true) {
                            v20 = v20
                            v24 := GenS1{F0: 65535}
                            if (1 != 42) {
                                var v25 int32 = 559
                                _ = v25
                            }
                            _ = v24
                        }
                        _ = fi21
                    }
                    fmt.Println(p0, fi10, fi10)
                    _ = v20
                }
                var v26 int8 = -1
                for fi27 := 0; fi27 < 15; fi27++ {
                    fmt.Println(v26, p0)
                    _ = fi27
                }
                _ = fi10
                _ = v26
            }
        }
    } else {
        var v28 uint16 = 32767
        p0 -= 128
        _ = v28
    }
    v29 := GenS2{F0: -10}
    v30 := GenS2{F0: 1}
    _ = p0
    _ = v29
    _ = v30
    return 42
}
func genFn1(p0 uint64) bool {
    fmt.Println(p0)
    for fi31 := 0; fi31 < 7; fi31++ {
        v32 := GenS2{F0: 0}
        if true {
            switch (fi31 | fi31) {
            case 128:
                if (!false) {
                    var v33 uint = 400
                    for fi34 := 0; fi34 < 13; fi34++ {
                        var v35 float64 = (1e10 - 3.14159)
                        v36 := []int{(fi31 | (fi31 / ((fi34) | 1))), (^(fi31)), 127, 127}
                        v37 := GenS3{F0: -694, F1: 2}
                        _ = fi34
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                    if (!true) {
                        v38 := []int{}
                        _ = v38
                    } else {
                        var v39 uint16 = 32767
                        v40 := []int{fi31}
                        _ = v39
                        _ = v40
                    }
                    _ = v33
                } else {
                    v41 := []int{9223372036854775807, fi31}
                    _ = v41
                }
            case 19:
                v42 := []int{(fi31 / ((fi31) | 1)), fi31}
                _ = v42
            default:
                var v43 uint64 = p0
                _ = v43
            }
            v44 := []int{fi31, (^(fi31))}
            _ = v44
        } else {
            v45 := []int{32767}
            _ = v45
        }
        var v46 string = ("\n\t" + "hello")
        _ = fi31
        _ = v32
        _ = v46
    }
    var v47 uint16 = 553
    _ = p0
    _ = v47
    return true
}

func main() {
    v1 := []string{}
    var v2 bool = false
    v3 := make(map[uint32]bool)
    var v4 int8 = -84
    var v5 string = "a"
    var v48 int = 32768
    v49 := GenS1{F0: 92}
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    if true {
        for fi50 := 0; fi50 < 13; fi50++ {
            v51 := GenS1{F0: 2147483648}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi50
            _ = v51
        }
        v52 := GenS3{F0: 447, F1: 42}
        _ = v52
    } else {
        var v53 uint32 = 963
        var v54 int32 = 42
        _ = v53
        _ = v54
    }
    v3[79] = (v5 < (v5 + v5))
    v48 = 128
    delete(v3, 2147483647)
    delete(v3, 2147483647)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
}
